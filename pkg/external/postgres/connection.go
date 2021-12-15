package postgres

import "github.com/andreashanson/golang-rabbitmq/pkg/config"

func NewConnection(c *config.PostgresConfig) *PostgresConnection {
	return &PostgresConnection{}
}

type PostgresConnection struct {
}
