package queue

import (
	"fmt"
	"rabbitmq/email"

	"github.com/streadway/amqp"
)

func Consumer() {
	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msg, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Succesfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for message")

	<-forever

}

func ConsumerEmail() {
	fmt.Println("Consumer Email")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msg, err := ch.Consume(
		"EmailQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			email.EmailV2(d.Body)
		}
	}()

	fmt.Println("Succesfully connected to our RabbitMQ instance")
	fmt.Println(" [*] - waiting for message")

	<-forever

}
