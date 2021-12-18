package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
	"github.com/streadway/amqp"
)

type Repository interface {
	Consume(queue string) (<-chan amqp.Delivery, error)
	Ack(tag uint64) error
	Nack(tag uint64) error
	Publish(b []byte, queue string) error
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Consume(queue string) (<-chan amqp.Delivery, error) {
	msg, err := s.repo.Consume(queue)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func (s *Service) HandleMessages(msgType string, msgs <-chan amqp.Delivery, out chan<- msg.Message, errs chan<- error) {
	switch msgType {
	case "scheduler":
		for m := range msgs {
			fmt.Println(string(m.Body))
		}
	case "jobs":
		for m := range msgs {
			var mb msg.Body
			err := json.Unmarshal(m.Body, &mb)
			if err != nil {
				s.repo.Nack(m.DeliveryTag)
				//s.repo.Ack(m.DeliveryTag)
				errs <- err
				continue
			}
			message := msg.Message{
				DeliveryTag: m.DeliveryTag,
				Body:        mb,
				Exchange:    m.Exchange,
			}
			err = s.repo.Ack(m.DeliveryTag)
			if err != nil {
				fmt.Println("Could not ack msg")
				errs <- err
				continue
			}
			out <- message
			continue
		}
	}
}
