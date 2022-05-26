package adapter

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-redis/redis/v8"
	"gitlab.com/go-prism/go-rbac-proxy/pkg/rbac"
)

type RedisAdapter struct {
	client redis.UniversalClient
}

func NewRedisAdapter(ctx context.Context, opts *redis.UniversalOptions) *RedisAdapter {
	log := logr.FromContextOrDiscard(ctx).WithName("redis")
	log.V(2).Info("attempting to connect to redis")
	log.V(3).Info("using redis client configuration", "Options", opts)
	// open connection
	return &RedisAdapter{
		client: redis.NewUniversalClient(opts),
	}
}

func (r *RedisAdapter) SubjectHasGlobalRole(ctx context.Context, subject, role string) (bool, error) {
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

func getKey(subject, resource string, action *rbac.Verb) string {
	if action == nil {
		return fmt.Sprintf("%s-%s", subject, resource)
	}
	return fmt.Sprintf("%s-%s-%s", subject, resource, action)
}
