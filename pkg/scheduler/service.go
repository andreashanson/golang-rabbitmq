package scheduler

import (
	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
)

type Service struct {
	Jobs     []job
	Producer *producer.Service
}

func NewService(p *producer.Service) *Service {
	jobs := createJobs()
	return &Service{
		Jobs:     *jobs,
		Producer: p,
	}
}

func (s *Service) Start(errChan chan error) {
	for _, job := range s.Jobs {
		job.startJob(s.Producer)
	}
}
