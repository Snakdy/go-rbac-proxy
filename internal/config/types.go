package config

import "github.com/go-redis/redis/v8"

type Globals map[string][]string

type Configuration struct {
	Globals Globals `yaml:"globals"`
	Adapter Adapter `yaml:"adapter"`
}

type Adapter struct {
	Mode     string                 `yaml:"mode"`
	Redis    redis.UniversalOptions `yaml:"redis"`
	Postgres PostgresAdapter        `yaml:"postgres"`
}

type PostgresAdapter struct {
	DSN string `yaml:"dsn"`
}
