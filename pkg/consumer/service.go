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

func (s *Service) HandleMessages(msgType string, msgs <-chan amqp.Delivery, out chan msg.Message) error {
	switch msgType {
	case "jobs":
		fmt.Println("jobs")
		for m := range msgs {
			fmt.Println(string(m.Body))
		}
	case "scheduler":
		fmt.Println("sched")

		for m := range msgs {
			var mb msg.Body
			err := json.Unmarshal(m.Body, &mb)
			if err != nil {
				fmt.Println("Could not unmarshall msg", m.DeliveryTag)
				s.repo.Nack(m.DeliveryTag)
				//s.repo.Ack(m.DeliveryTag)
				return err
			}

			message := msg.Message{
				DeliveryTag: m.DeliveryTag,
				Body:        mb,
				Exchange:    m.Exchange,
			}

			out <- message
			err = s.repo.Ack(m.DeliveryTag)
			if err != nil {
				fmt.Println("Could not ack msg")
				return err
			}
		}
	}
	return nil
}
