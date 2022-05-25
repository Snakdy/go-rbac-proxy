package config

import (
	"context"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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
	return &c, nil
}
