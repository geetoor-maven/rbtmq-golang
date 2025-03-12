package main

import (
	"github.com/rabbitmq/amqp091-go"
)

// ConnectToRabbitMQ : fungsi untuk koneksi ke rabbit mq
func ConnectToRabbitMQ() (*amqp091.Connection, *amqp091.Channel) {
	con, errCon := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(errCon, "Failed to connect to RabbitMQ")
	ch, errChan := con.Channel()
	FailOnError(errChan, "Failed to open a channel")
	return con, ch
}
