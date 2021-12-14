package scheduler

import (
	"time"

	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
	"github.com/robfig/cron"
)

type job struct {
	name         string
	cronJob      *cron.Cron
	cronSchedule string
}

func createJobs() *[]job {
	c := cron.New()
	return &[]job{
		{
			name:         "google",
			cronJob:      c,
			cronSchedule: "@every 10s",
		},
		{
			name:         "slack",
			cronJob:      c,
			cronSchedule: "@every 1m",
		},
		{
			name:         "linkedin",
			cronJob:      c,
			cronSchedule: "@every 10",
		},
		{
			name:         "facebook",
			cronJob:      c,
			cronSchedule: "@every 5s",
		},
		{
			name:         "instagram",
			cronJob:      c,
			cronSchedule: "@every 10m",
		},
	}
}

func (j job) startJob(p producer.Service) error {
	var err error
	err = j.cronJob.AddFunc(j.cronSchedule, func() {
		start_time := time.Now().UTC()
		err2 := p.Publish([]byte(`{"type":"`+j.name+`", "start_time":"`+start_time.String()+`"}`), "scheduler")
		if err2 != nil {
			err = err2
		}
	})
	if err != nil {
		return err
	}
	j.cronJob.Start()
	return nil
}
