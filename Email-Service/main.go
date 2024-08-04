package main

import (
	"Go-Nats-Demo/common"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Email Service is Up & Running")
	// Connect to a NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	messages := make(chan *nats.Msg)

	for _, subject := range common.GetAllSubjects() {
		_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
			messages <- msg
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	go func() {
		for msg := range messages {
			log.Printf("Email Service: %s\n", string(msg.Data))
		}
	}()

	// Keep the program running to receive messages indefinitely
	select {}

}
