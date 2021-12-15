package rabbit

import (
	"github.com/andreashanson/golang-rabbitmq/pkg/config"
	"github.com/streadway/amqp"
)

func Connect(cfg *config.RabbitMQConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":5672/")
	return conn, err
}
