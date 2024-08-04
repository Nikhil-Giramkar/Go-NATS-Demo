package main

import (
	"Go-Nats-Demo/common"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("Order Service is Up & Running")
	// Connect to a NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	fmt.Println()
	fmt.Println("Order Placed By Customer")
	nc.Publish(common.ORDER_PLACED, []byte("Bluetooth Speakers Ordered"))
	
	fmt.Println()
	time.Sleep(1 * time.Second)
	fmt.Println("Order Packed")
	nc.Publish(common.ORDER_PACKED, []byte("Order Packed"))

	fmt.Println()
	time.Sleep(2 * time.Second)
	fmt.Println("Delivery Agent Assigned")
	nc.Publish(common.DELIVERY_STARTED, []byte("Delivery Started - Agent Assigned"))

	fmt.Println()
	time.Sleep(2 * time.Second)
	fmt.Println("Delivery Complete")
	nc.Publish(common.DELIVERY_STARTED, []byte("Delivery Complete"))

	fmt.Println()
	time.Sleep(1 * time.Second)
	fmt.Println("Payment Success")
	nc.Publish(common.PAYMENT_SUCCESS, []byte("Payment Successful"))
}
