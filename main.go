package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Print("golang rabbit mq")

	http.HandleFunc("/checkout", CheckOut)
	err := http.ListenAndServe(":8080", nil)
	FailOnError(err, "error starting http server")
}
