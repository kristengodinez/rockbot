package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(CreditCardValidatorServer)
	log.Fatal(http.ListenAndServe(":5001", handler))
}
