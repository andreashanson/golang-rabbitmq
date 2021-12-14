package rabbit

import (
	"github.com/streadway/amqp"
)

type RabbitMQRepo struct {
	Channel        *amqp.Channel
	schedulerQueue *amqp.Queue
}

func NewRabbitMQRepo(conn *amqp.Connection) (RabbitMQRepo, error) {
	ch, err := conn.Channel()

	if err != nil {
		return RabbitMQRepo{}, err
	}

	schedulerQueue, err := ch.QueueDeclare(
		"scheduler", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		return RabbitMQRepo{}, err
	}
	return RabbitMQRepo{Channel: ch, schedulerQueue: &schedulerQueue}, nil
}

func (r *RabbitMQRepo) Publish(b []byte, channel string) error {

	return r.Channel.Publish(
		"",
		channel,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
}

func (r *RabbitMQRepo) Consume(channel string) (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		channel,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (r *RabbitMQRepo) Ack(tag uint64) error {
	return r.Channel.Ack(tag, true)
}

func (r *RabbitMQRepo) Nack(tag uint64) error {
	return r.Channel.Nack(tag, true, false)
}
