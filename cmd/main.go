package main

import (
	"fmt"
	"log"
	"time"

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

	//for i := 0; i < 100; i++ {
	//	indexString := strconv.Itoa(i)
	//	body := "{'hello':'world', 'index':'" + indexString + "'}"
	//
	//	err := prd.Publish([]byte(body))
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	fmt.Println("Sleep for five seconds before consume msgs")
	time.Sleep(5 * time.Second)
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
			fmt.Println(msg)
		//case forever := <-forever:
		//	fmt.Print("DIE")
		//	fmt.Println(forever)
		case errs := <-errChan:
			fmt.Println(errs)
		case errs2 := <-errChan2:
			fmt.Println(errs2)
		}
	}
}
