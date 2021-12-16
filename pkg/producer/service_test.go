package producer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		r Repository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "Test create new service",
			args: args{r: mockRepository{}},
			want: &Service{repo: mockRepository{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.r)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestService_Publish(t *testing.T) {
	type args struct {
		b     []byte
		queue string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			name:    "Happy publish message",
			s:       &Service{repo: mockRepository{err: nil}},
			args:    args{b: []byte("test"), queue: "test"},
			wantErr: false},
		{
			name:    "Not happy publish message",
			s:       &Service{repo: mockRepository{err: errors.New("Fake error")}},
			args:    args{b: []byte("test"), queue: "test"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Publish(tt.args.b, tt.args.queue); (err != nil) != tt.wantErr {
				t.Errorf("Service.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type mockRepository struct {
	err error
}

func (mr mockRepository) Publish(b []byte, queue string) error {
	return mr.err
}
