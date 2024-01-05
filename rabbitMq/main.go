package main

import (
	"github.com/streadway/amqp"
	"log"
)

// req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
// url := "https://in.bookmyshow.com/sports/india-vs-new-zealand-icc-mens-cwc-2023/seat-layout/aerialcanvas/ET00367570/HCCD/10013?groupEventCode=ET00367219"
func main() {
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

	queueName := "first_direct_queue"
	msgs, err := ch.Consume(
		queueName,
		"consumer1",
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
