package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {

	//docker run -d --name rabbitmq-container -p 5672:5672 -p 15672:15672 rabbitmq:management
	//guest:guest http://localhost:15672

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	queueName := "test_queue"
	msgs, err := ch.Consume(
		queueName,
		"",
		true,  // Auto-acknowledgment
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Args
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %v", err)
	}

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
	}
}
