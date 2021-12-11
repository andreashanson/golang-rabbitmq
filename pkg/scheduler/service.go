package scheduler

import (
	"fmt"
	"time"

	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
	"github.com/robfig/cron"
)

type Service struct {
	Timer    *time.Ticker
	Jobs     *cron.Cron
	Producer producer.Service
}

func NewService(t int64, p producer.Service) *Service {
	c := cron.New()

	tt := time.Duration(t)
	timer := time.NewTicker(tt * time.Second)

	return &Service{
		Timer:    timer,
		Jobs:     c,
		Producer: p,
	}
}

func (s *Service) Start(errChan chan error) <-chan time.Time {
	fmt.Println("Start schduler service.")

	s.Jobs.AddFunc("*/5 * * * *", func() {
		t := time.Now().UTC()
		t = t.Truncate(time.Duration(t.Hour()) * time.Hour)
		err := s.Producer.Publish([]byte(`{"type":"hotspot", "msg":"Hello", "start_time":"` + t.String() + `"}`))
		if err != nil {
			fmt.Println("Could not Publish msg")
			errChan <- err
		}
	})
	s.Jobs.Start()
	return s.Timer.C
}
