package scheduler

import (
	"fmt"
	"time"

	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
	"github.com/robfig/cron"
)

type Service struct {
	Timer    *time.Ticker
	CronJob  *cron.Cron
	Producer producer.Service
	Funcs    func()
	Jobs     []job
}

func NewService(t int64, p producer.Service) *Service {
	c := cron.New()

	tt := time.Duration(t)
	timer := time.NewTicker(tt * time.Second)

	return &Service{
		Timer:    timer,
		CronJob:  c,
		Producer: p,
		Funcs:    getGoogleData,
		Jobs:     *getJobs(),
	}
}

func (s *Service) Start(errChan chan error) <-chan time.Time {
	fmt.Println("Start schduler service.")

	for _, job := range s.Jobs {
		err := s.CronJob.AddFunc(job.cronSchedule, job.cronFunc)
		if err != nil {
			fmt.Println(err)
		}
	}

	s.CronJob.Start()
	return s.Timer.C
}
