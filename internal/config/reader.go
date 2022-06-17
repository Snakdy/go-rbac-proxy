package config

import (
	"context"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func Read(ctx context.Context, path string) (*Configuration, error) {
	log := logr.FromContextOrDiscard(ctx).WithValues("Path", path)
	log.V(1).Info("reading configuration file")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Error(err, "failed to read file")
		return nil, err
	}
	log.V(2).Info("successfully read file", "Raw", string(data))
	// parse yaml
	var c Configuration
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Error(err, "failed to unmarshal yaml")
		return nil, err
	}
	log.V(1).Info("successfully unmarshalled yaml")
	// expand environment variables in
	// secret config items (see #1)
	c.Adapter.Redis.Password = os.ExpandEnv(c.Adapter.Redis.Password)
	c.Adapter.Redis.SentinelPassword = os.ExpandEnv(c.Adapter.Redis.SentinelPassword)
	c.Adapter.Postgres.DSN = os.ExpandEnv(c.Adapter.Postgres.DSN)

	return &c, nil
}
