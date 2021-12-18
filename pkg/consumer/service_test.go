package consumer

import (
	"errors"
	"testing"

	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	type args struct {
		r Repository
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		{
			name: "Test New service",
			args: args{r: mockRepository{deliverChan: make(chan amqp.Delivery), err: nil}},
			want: &Service{
				repo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotNil(t, NewService(tt.args.r))
		})
	}
}

func TestService_Consume(t *testing.T) {

	deliveryChan := make(chan amqp.Delivery)

	type args struct {
		queue string
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		want    <-chan amqp.Delivery
		wantErr bool
	}{
		{
			name:    "Test consume happy path",
			s:       &Service{repo: mockRepository{deliverChan: deliveryChan, err: nil}},
			args:    args{queue: "test"},
			want:    deliveryChan,
			wantErr: false,
		},
		{
			name:    "Test consume not happy path",
			s:       &Service{repo: mockRepository{deliverChan: deliveryChan, err: errors.New("fake error")}},
			args:    args{queue: "test"},
			want:    deliveryChan,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Consume(tt.args.queue)
			if err != nil {
				assert.Error(t, err)
				assert.True(t, tt.wantErr == true)
				return
			}
			assert.Equal(t, got, tt.want)
			assert.True(t, tt.wantErr == false)
		})
	}
}

func TestService_HandleMessages(t *testing.T) {
	type args struct {
		msgType    string
		deliveries []amqp.Delivery
	}
	tests := []struct {
		name      string
		s         *Service
		args      args
		msgsCount int
		errsCount int
	}{
		{
			name: "Test handlemessages happy path",
			s:    &Service{repo: mockRepository{}},
			args: args{msgType: "jobs", deliveries: []amqp.Delivery{
				{
					DeliveryTag: 1,
					Body:        []byte(`{"type":"google", "start_time":"2021-12-24"}`),
				},
				{
					DeliveryTag: 2,
					Body:        []byte(`{"type":"facebook", "start_time":"2021-12-24"}`),
				},
			}},
			msgsCount: 2,
			errsCount: 0,
		},
		{
			name: "Test handlemessages happy path with errors",
			s:    &Service{repo: mockRepository{}},
			args: args{msgType: "jobs", deliveries: []amqp.Delivery{
				{
					DeliveryTag: 1,
					Body:        []byte(`{"type":"google", "start_time":"2021-12-24"}`),
				},
				{
					DeliveryTag: 2,
					Body:        []byte(`not json will give error`),
				},
			}},
			msgsCount: 1,
			errsCount: 1,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			deliveries := make(chan amqp.Delivery)
			outChan := make(chan msg.Message)
			errChan := make(chan error)
			go tt.s.HandleMessages(tt.args.msgType, deliveries, outChan, errChan)

			var msgs []msg.Message
			var errs []error

			for _, d := range tt.args.deliveries {
				deliveries <- d

				select {
				case m := <-outChan:
					assert.True(t, len(m.Body.Type) > 0)
					msgs = append(msgs, m)
				case e := <-errChan:
					errs = append(errs, e)

				}
			}
			assert.Equal(t, tt.msgsCount, len(msgs))
			assert.Equal(t, tt.errsCount, len(errs))
		})
	}
}

type mockRepository struct {
	deliverChan <-chan amqp.Delivery
	err         error
}

func (mr mockRepository) Consume(queue string) (<-chan amqp.Delivery, error) {
	return mr.deliverChan, mr.err
}
func (mr mockRepository) Ack(tag uint64) error {
	return mr.err
}
func (mr mockRepository) Nack(tag uint64) error {
	return mr.err
}
func (mr mockRepository) Publish(b []byte, queue string) error {
	return nil
}
