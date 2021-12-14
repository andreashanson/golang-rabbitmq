package config

import "os"

type Config struct {
	RabbitMQ *RabbitMQConfig
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
	}
}
