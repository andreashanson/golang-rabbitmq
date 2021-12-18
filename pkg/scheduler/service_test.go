package scheduler

import (
	"reflect"
	"testing"

	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
)

func TestNewService(t *testing.T) {
	type args struct {
		p *producer.Service
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Start(t *testing.T) {
	type args struct {
		errChan chan error
	}
	tests := []struct {
		name string
		s    *Service
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		//errChan := make(chan error)
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Start(tt.args.errChan)
		})
	}
}
