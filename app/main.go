package main

import (
	"log"
	"net/http"
)

type InMemoryCreditCardStore struct{}

func (i *InMemoryCreditCardStore) GetCardValidation(number string) bool {
	return true
}

func main() {
	server := &CreditCardValidatorServer{&InMemoryCreditCardStore{}}
	log.Fatal(http.ListenAndServe(":5001", server))
}
