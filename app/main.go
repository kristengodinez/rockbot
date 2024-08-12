package main

import (
	"log"
	"net/http"
)

func main() {
	server := &CreditCardValidatorServer{&InMemoryCreditCardStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
