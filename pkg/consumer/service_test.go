package consumer

import (
	"errors"
	"fmt"
	"testing"
	"time"

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
			}
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestService_HandleMessages(t *testing.T) {
	type args struct {
		msgType    string
		deliveries []amqp.Delivery
	}
	tests := []struct {
		name    string
		s       *Service
		args    args
		wantErr bool
	}{
		{
			name: "Test handlemessages happy path",
			s:    &Service{repo: mockRepository{}},
			args: args{msgType: "jobs", deliveries: []amqp.Delivery{
				{
					Acknowledger:    nil,
					Headers:         map[string]interface{}{},
					ContentType:     "",
					ContentEncoding: "",
					DeliveryMode:    0,
					Priority:        0,
					CorrelationId:   "",
					ReplyTo:         "",
					Expiration:      "",
					MessageId:       "",
					Timestamp:       time.Time{},
					Type:            "",
					UserId:          "",
					AppId:           "",
					ConsumerTag:     "",
					MessageCount:    0,
					DeliveryTag:     1,
					Redelivered:     false,
					Exchange:        "",
					RoutingKey:      "",
					Body:            []byte(`{"type":"google", "start_time":"2021-12-24"}`),
				},
				{
					Acknowledger:    nil,
					Headers:         map[string]interface{}{},
					ContentType:     "",
					ContentEncoding: "",
					DeliveryMode:    0,
					Priority:        0,
					CorrelationId:   "",
					ReplyTo:         "",
					Expiration:      "",
					MessageId:       "",
					Timestamp:       time.Time{},
					Type:            "",
					UserId:          "",
					AppId:           "",
					ConsumerTag:     "",
					MessageCount:    0,
					DeliveryTag:     2,
					Redelivered:     false,
					Exchange:        "",
					RoutingKey:      "",
					Body:            []byte(`{"type":"facebook", "start_time":"2021-12-24"}`),
				}}},
			wantErr: false,
		},
		//{
		//	name:    "Test handlemessages not happy path",
		//	s:       &Service{repo: mockRepository{}},
		//	args:    args{msgType: "jobs ", deliveries: deliveryChan, out: outChan},
		//	wantErr: true,
		//},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			deliveries := make(chan amqp.Delivery)
			outChan := make(chan msg.Message)
			errChan := make(chan error)
			go func() {
				tt.s.HandleMessages(tt.args.msgType, deliveries, outChan, errChan)
			}()

			var msgs []msg.Message

			for _, d := range tt.args.deliveries {
				deliveries <- d

				select {
				case msg := <-outChan:
					msgs = append(msgs, msg)
				case err := <-errChan:
					fmt.Println(err)
				}
			}
			assert.Equal(t, len(msgs), 2)

		})

		//t.Run(tt.name, func(t *testing.T) {
		//	err2 := tt.s.HandleMessages(tt.args.msgType, tt.args.deliveries, tt.args.out)
		//	if err2 != nil {
		//		t.Errorf("Service.HandleMessages() error2 = %v, wantErr %v", err2, tt.wantErr)
		//	}
		//})
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
