package main

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"time"
)

// processSendEmail : fungsi ini hanya simulasi pengiriman email.
func processSendEmail(order Order) {
	fmt.Printf("Proses pengiriman email konfirmasi ke %s ", order.Email)
	fmt.Println()
	time.Sleep(2 * time.Second) // berikan delay pengiriman
	fmt.Println("Email berhasil di kirim")
}

func GetMessageFromQueue(channel *amqp091.Channel, queueName string) {

	queueDeclare, errQueueDec := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	FailOnError(errQueueDec, "Failed to declare a queue")

	msgs, errConsume := channel.Consume(
		queueDeclare.Name,
		"",
		true,  // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,
	)
	FailOnError(errConsume, "Failed to consume a message")

	for {
		select {
		case msg := <-msgs:
			var order Order
			err := json.Unmarshal(msg.Body, &order)
			FailOnError(err, "Failed to unmarshal JSON")
			processSendEmail(order)
		case <-time.After(10 * time.Second): // jika tidak terdapat data dalam antrian selama 10 detik
			fmt.Println("Tidak ada order baru dalam 10 detik, menunggu kembali...")
		}
	}
}
