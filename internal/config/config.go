package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Listen         string `envconfig:"LISTEN"`
	ListenInternal string `envconfig:"LISTEN_INTERNAL"`

	Postgres Postgres
}

type Postgres struct {
	Host           string `envconfig:"POSTGRES_HOST"`
	User           string `envconfig:"POSTGRES_USER"`
	Password       string `envconfig:"POSTGRES_PASSWORD"`
	DB             string `envconfig:"POSTGRES_DB"`
	SimpleProtocol bool   `envconfig:"POSTGRES_SIMPLE_PROTOCOL"`
}

func New() *Config {
	if err := godotenv.Load(".env"); err != nil {
		return nil
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil
	}
	return &cfg
}
