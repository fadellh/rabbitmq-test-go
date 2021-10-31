package main

import (
	"fmt"
	"rabbitmq/queue"
	"time"
)

func main() {
	fmt.Println("Golang email app running")
	// email.Email()
	// email.EmailV2()
	email := "fadel.lukman.dev@gmail.com"
	queue.PubliserEmail(email)

	time.Sleep(8 * time.Second)

	queue.ConsumerEmail()

}
