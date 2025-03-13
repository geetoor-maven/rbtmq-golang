package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ValidateHttpMethod(w http.ResponseWriter, r *http.Request, httpMehtod string) {
	if r.Method != httpMehtod {
		http.Error(w, "only method "+httpMehtod+" is allowed", http.StatusMethodNotAllowed)
		return
	}
}

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	FailOnError(err, "Error decoding body")
}
