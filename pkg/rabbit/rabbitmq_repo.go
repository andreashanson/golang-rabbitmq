package rabbit

import (
	"github.com/streadway/amqp"
)

type RabbitMQRepo struct {
	Channel  *amqp.Channel
	queue    *amqp.Queue
	ChanName string
}

func NewRabbitMQRepo(conn *amqp.Connection) (RabbitMQRepo, error) {
	ch, err := conn.Channel()

	if err != nil {
		return RabbitMQRepo{}, err
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return RabbitMQRepo{}, err
	}
	return RabbitMQRepo{Channel: ch, queue: &q, ChanName: q.Name}, nil
}

func (r *RabbitMQRepo) Publish(b []byte) error {

	return r.Channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		},
	)
}

func (r *RabbitMQRepo) Consume() (<-chan amqp.Delivery, error) {
	return r.Channel.Consume(
		r.queue.Name,
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
