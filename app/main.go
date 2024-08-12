package main

import (
	"log"
	"net/http"
)

func main() {
	handler := &CreditCardValidatorServer{}
	log.Fatal(http.ListenAndServe(":5001", handler))
}
