package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("golang rabbit mq")
	fmt.Println("---------------------")

	// mari kita gunakan go routine dari golang agar proses pengiriman pesan dapat async
	go func() {
		connectToRabbitMQ, channel := ConnectToRabbitMQ()
		defer connectToRabbitMQ.Close()
		defer channel.Close()

		GetMessageFromQueue(channel, "email_queue")
	}()

	fmt.Println("Running api checkout")
	fmt.Println("---------------------")
	http.HandleFunc("/checkout", CheckOut)
	err := http.ListenAndServe(":8080", nil)
	FailOnError(err, "error starting http server")
}
