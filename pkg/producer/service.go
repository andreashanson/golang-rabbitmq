package producer

import (
	"encoding/json"

	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
)

type Repository interface {
	Publish(b []byte, queue string) error
}

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Publish(b []byte, queue string) error {
	var body msg.Body
	err := json.Unmarshal(b, &body)
	if err != nil {
		return err
	}
	return s.repo.Publish(b, queue)
}
