package adapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"net/url"
	"strings"
)

type RedisAdapter struct {
	Unimplemented
	client redis.UniversalClient
}

func NewRedisAdapter(ctx context.Context, opts *redis.UniversalOptions) *RedisAdapter {
	log := logr.FromContextOrDiscard(ctx).WithName("redis")
	log.V(2).Info("attempting to connect to redis")
	log.V(3).Info("using redis client configuration", "Options", opts)
	// open connection
	rdb := redis.NewUniversalClient(opts)
	// add OpenTelemetry tracing
	rdb.AddHook(redisotel.NewTracingHook())
	return &RedisAdapter{
		client: rdb,
	}
}

func (r *RedisAdapter) SubjectHasGlobalRole(ctx context.Context, subject, role string) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_subjectHasGlobalRole", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role", role).WithName("redis")
	log.V(1).Info("checking if subject has global role")
	key := getKey(subject, role, nil)
	log.V(2).Info("getting redis key", "Key", key)
	// fetch the value from redis
	val, err := r.client.Get(ctx, key).Int()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to retrieve global role binding")
		return false, err
	}
	log.V(1).Info("retrieved role information", "Key", key, "Value", val)
	return val > 0, nil
}

func (r *RedisAdapter) SubjectCanDoAction(ctx context.Context, subject, resource string, action rbac.Verb) (bool, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_subjectCanDoAction", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("resource", resource),
		attribute.String("action", action.String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("redis")
	log.V(1).Info("checking if subject has role")
	key := getKey(subject, resource, &action)
	sudoAction := rbac.Verb_SUDO
	keySudo := getKey(subject, resource, &sudoAction)
	log.V(2).Info("getting redis keys", "Keys", []string{key, keySudo})
	// fetch the value from redis
	values, err := r.client.MGet(ctx, key, keySudo).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.V(1).Info("matching role could not be found")
			return false, nil
		}
		log.Error(err, "failed to retrieve role binding")
		return false, err
	}
	log.V(1).Info("retrieved role information", "Key", key, "Values", values)
	// iterate the values
	for _, v := range values {
		// try to convert it to an int
		val, ok := v.(string)
		// if it's an int and greater than 0
		// we have a match
		if ok && val == "1" {
			return true, nil
		}
	}
	return false, nil
}

func (r *RedisAdapter) Add(ctx context.Context, subject, resource string, action rbac.Verb) error {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_add", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("resource", resource),
		attribute.String("action", action.String()),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Resource", resource, "Action", action.String()).WithName("redis")
	log.V(1).Info("creating role binding")
	key := getKey(subject, resource, &action)
	log.V(2).Info("setting redis key", "Key", key, "Value", 1)
	// set the value in redis
	if err := r.client.SetNX(ctx, key, 1, 0).Err(); err != nil {
		log.Error(err, "failed to add role binding")
		return err
	}
	return nil
}

func (r *RedisAdapter) AddGlobal(ctx context.Context, subject, role string) error {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_addGlobal", trace.WithAttributes(
		attribute.String("subject", subject),
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject, "Role", role).WithName("redis")
	log.V(1).Info("creating global role binding")
	key := getKey(subject, role, nil)
	log.V(2).Info("setting redis key", "Key", key, "Value", 1)
	// set the value in redis
	if err := r.client.SetNX(ctx, key, 1, 0).Err(); err != nil {
		log.Error(err, "failed to add global role binding")
		return err
	}
	return nil
}

func (r *RedisAdapter) ListBySub(ctx context.Context, subject string) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_listBySub", trace.WithAttributes(
		attribute.String("subject", subject),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Subject", subject).WithName("redis")
	log.V(1).Info("scanning for roles with subject")

	var items []*rbac.RoleBinding

	match := fmt.Sprintf("%s/*", subject)
	log.V(2).Info("scanning for keys", "Match", match)
	iter := r.client.Scan(ctx, 0, match, 0).Iterator()
	for iter.Next(ctx) {
		// do something
		items = append(items, parseRoleBinding(iter.Val()))
	}
	if err := iter.Err(); err != nil {
		log.Error(err, "failed to list roles by subject")
		return nil, err
	}
	log.Info("successfully fetched roles for subject", "Count", len(items))
	return items, nil
}

func (r *RedisAdapter) ListByRole(ctx context.Context, role string) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_listByRole", trace.WithAttributes(
		attribute.String("role", role),
	))
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithValues("Role", role).WithName("redis")
	log.V(1).Info("scanning for bindings with role")

	var items []*rbac.RoleBinding

	iter := r.client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		// do something
		rb := parseRoleBinding(iter.Val())
		if rb.GetResource() != role {
			continue
		}
		items = append(items, rb)
	}
	if err := iter.Err(); err != nil {
		log.Error(err, "failed to list roles by subject")
		return nil, err
	}
	log.Info("successfully fetched bindings for role", "Count", len(items))
	return items, nil
}

func (r *RedisAdapter) List(ctx context.Context) ([]*rbac.RoleBinding, error) {
	ctx, span := otel.Tracer("").Start(ctx, "adapter_redis_list")
	defer span.End()
	log := logr.FromContextOrDiscard(ctx).WithName("redis")
	log.V(1).Info("scanning for bindings")

	var items []*rbac.RoleBinding

	iter := r.client.Scan(ctx, 0, "*", 0).Iterator()
	for iter.Next(ctx) {
		// do something
		rb := parseRoleBinding(iter.Val())
		items = append(items, rb)
	}
	if err := iter.Err(); err != nil {
		log.Error(err, "failed to list roles")
		return nil, err
	}
	log.Info("successfully fetched bindings", "Count", len(items))
	return items, nil
}

func parseRoleBinding(s string) *rbac.RoleBinding {
	var verb rbac.Verb
	// split the value
	bits := strings.Split(s, "/")
	// there will always be 2 chunks, so grab them first
	sub := mustDecode(bits[0])
	resource := mustDecode(bits[1])
	// if there's a 3rd chunk, use that as the verb
	if len(bits) > 2 {
		verb = rbac.Verb(rbac.Verb_value[bits[2]])
	}
	return &rbac.RoleBinding{
		Subject:  sub,
		Resource: resource,
		Action:   verb,
	}
}

func mustDecode(s string) string {
	r, err := url.PathUnescape(s)
	if err != nil {
		return ""
	}
	return r
}

func getKey(subject, resource string, action *rbac.Verb) string {
	// encode so we can store
	// values with slash
	subject = url.PathEscape(subject)
	resource = url.PathEscape(resource)
	if action == nil {
		return fmt.Sprintf("%s/%s", subject, resource)
	}
	return fmt.Sprintf("%s/%s/%s", subject, resource, action)
}
