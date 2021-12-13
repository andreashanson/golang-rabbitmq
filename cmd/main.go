package main

import (
	"fmt"
	"log"

	"github.com/andreashanson/golang-rabbitmq/pkg/config"
	"github.com/andreashanson/golang-rabbitmq/pkg/consumer"
	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
	"github.com/andreashanson/golang-rabbitmq/pkg/rabbit"
	"github.com/andreashanson/golang-rabbitmq/pkg/scheduler"
)

func main() {

	cfg := config.Config{
		RabbitMQ: &config.RabbitMQConfig{
			Host:     "localhost",
			User:     "guest",
			Password: "guest",
		},
	}

	connection, err := rabbit.Connect(cfg.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}

	rabbitRepo, err := rabbit.NewRabbitMQRepo(connection)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitRepo.Channel.Close()

	prd := producer.New(&rabbitRepo)

	sch := scheduler.NewService(1, *prd)
	errChan2 := make(chan error)
	go func() {
		time := sch.Start(errChan2)
		for t := range time {
			fmt.Println(t)
		}
	}()

	cns := consumer.NewService(&rabbitRepo)
	msgs, err := cns.Consume()
	if err != nil {
		log.Fatal(err)
	}
	out := make(chan msg.Message)
	errChan := make(chan error)

	go func() {
		err := cns.HandleMessages(msgs, out)
		if err != nil {
			errChan <- err
		}
	}()

	for {
		select {
		case msg := <-out:
			fmt.Println("Message from out chan in main select", msg)
		case errs := <-errChan:
			fmt.Println(errs)
		case errs2 := <-errChan2:
			fmt.Println(errs2)
		}
	}
}
