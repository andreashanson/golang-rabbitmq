package scheduler

import (
	"time"

	"github.com/robfig/cron"
)

type job struct {
	name         string
	cronJob      *cron.Cron
	cronSchedule string
}

type prod interface {
	Publish(b []byte, queue string) error
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
			cronSchedule: "@every 3s",
		},
		{
			name:         "facebook",
			cronJob:      c,
			cronSchedule: "@every 5s",
		},
		{
			name:         "instagram",
			cronJob:      c,
			cronSchedule: "@every 4s",
		},
	}
}

func (j job) startJob(p prod) error {
	var err error
	err = j.cronJob.AddFunc(j.cronSchedule, func() {
		start_time := time.Now().UTC()
		err2 := p.Publish([]byte(`{"type":"`+j.name+`", "start_time":"`+start_time.String()+`"}`), "jobs")
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
