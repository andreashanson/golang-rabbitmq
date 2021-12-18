package scheduler

import (
	"errors"
	"testing"

	"github.com/robfig/cron"
	"github.com/stretchr/testify/assert"
)

func Test_job_startJob(t *testing.T) {
	c := cron.New()

	type args struct {
		p mockProducer
	}
	tests := []struct {
		name    string
		j       job
		args    args
		wantErr bool
	}{
		{
			name: "test publish happy path",
			j: job{
				name:         "google",
				cronJob:      c,
				cronSchedule: "@every 1s",
			},
			args:    args{mockProducer{err: nil}},
			wantErr: false,
		},
		{
			name: "test publish happy path",
			j: job{
				name:         "google",
				cronJob:      c,
				cronSchedule: "@every 1s",
			},
			args:    args{mockProducer{err: errors.New("FAKE ERROR")}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.j.startJob(tt.args.p)
			if err != nil {
				assert.Equal(t, tt.wantErr, true)
			}
			assert.Equal(t, tt.wantErr, false)
		})
	}
}

type mockProducer struct {
	err error
}

func (mp mockProducer) Publish(b []byte, queue string) error {
	return mp.err
}
