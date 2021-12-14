package rabbit

import (
	"github.com/streadway/amqp"
)

type RabbitMQRepo struct {
	Channel *amqp.Channel
}

func NewRabbitMQRepo(conn *amqp.Connection) (RabbitMQRepo, error) {
	ch, err := conn.Channel()

	if err != nil {
		return RabbitMQRepo{}, err
	}

	_, err = ch.QueueDeclare(
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

	_, err = ch.QueueDeclare(
		"jobs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)

	if err != nil {
		return RabbitMQRepo{}, err
	}
	return RabbitMQRepo{Channel: ch}, nil
}

func (r *RabbitMQRepo) Publish(b []byte, queue string) error {

	return r.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
}

func (r *RabbitMQRepo) Consume(queue string) (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		queue,
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
