package main

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"net/http"
)

type Order struct {
	OrderID string `json:"order_id"`
	Email   string `json:"email"`
	Status  string `json:"status"`
}

func publishMsgToQueue(channel *amqp091.Channel, queueName string, body []byte) {

	queueDeclare, errQueueDec := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	FailOnError(errQueueDec, "Failed to declare a queue")

	errPublish := channel.Publish(
		"",
		queueDeclare.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	FailOnError(errPublish, "Failed to publish a message")
}

func CheckOut(w http.ResponseWriter, r *http.Request) {
	ValidateHttpMethod(w, r, "POST")

	order := Order{}
	// read value from request body
	ReadFromRequestBody(r, &order)

	jsonBytes, errMarshal := json.Marshal(order)
	FailOnError(errMarshal, "Failed to marshal json")

	connectToRabbitMQ, channel := ConnectToRabbitMQ()
	defer connectToRabbitMQ.Close()
	defer channel.Close()

	// publish json to queue
	publishMsgToQueue(channel, "email_queue", jsonBytes)

	// write to response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"success", "message":"Order processed successfully", "order_id":"` + order.OrderID + `"}`))
}
