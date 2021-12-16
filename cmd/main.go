package main

import (
	"fmt"
	"log"

	"github.com/andreashanson/golang-rabbitmq/pkg/config"
	"github.com/andreashanson/golang-rabbitmq/pkg/consumer"
	"github.com/andreashanson/golang-rabbitmq/pkg/external/rabbit"
	"github.com/andreashanson/golang-rabbitmq/pkg/msg"
	"github.com/andreashanson/golang-rabbitmq/pkg/producer"
	"github.com/andreashanson/golang-rabbitmq/pkg/scheduler"
)

func main() {

	cfg := config.NewConfig()

	rabbitConnection, err := rabbit.Connect(cfg.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}

	rabbitRepo, err := rabbit.NewRabbitMQRepo(rabbitConnection)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitRepo.Channel.Close()

	prd := producer.New(&rabbitRepo)
	sch := scheduler.NewService(prd)
	errChan := make(chan error)
	out := make(chan msg.Message)

	go func() {
		sch.Start(errChan)
	}()

	jobConsumer := consumer.NewService(&rabbitRepo)
	jobMsgs, err := jobConsumer.Consume("jobs")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		jobConsumer.HandleMessages("jobs", jobMsgs, out, errChan)
	}()

	cns := consumer.NewService(&rabbitRepo)
	deliveries, err := cns.Consume("scheduler")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		cns.HandleMessages("scheduler", deliveries, out, errChan)
	}()

	for {
		select {
		case msg := <-out:
			fmt.Println("Message from out chan in main select", msg)
		case errs := <-errChan:
			fmt.Println("ERROR FROM ERR CHAN")
			fmt.Println(errs)
		}
	}
}
