package config

import "os"

type Config struct {
	RabbitMQ *RabbitMQConfig
	Postgres *PostgresConfig
}

type PostgresConfig struct {
	Host     string
	User     string
	Password string
}

type RabbitMQConfig struct {
	Host     string
	User     string
	Password string
}

func NewConfig() *Config {
	return &Config{
		RabbitMQ: &RabbitMQConfig{
			Host:     os.Getenv("RABBIT_HOST"),
			User:     os.Getenv("RABBIT_USER"),
			Password: os.Getenv("RABBIT_PW"),
		},
		Postgres: &PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PW"),
		},
	}
}
