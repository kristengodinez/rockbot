package main

func NewInMemoryCreditCardStore() *InMemoryCreditCardStore {
	return &InMemoryCreditCardStore{map[string]bool{}}
}

type InMemoryCreditCardStore struct {
	store map[string]bool
}

func (i *InMemoryCreditCardStore) GetCardValidation(creditCardNumber string) bool {
	i.store[creditCardNumber] = GetCardValidation(creditCardNumber)
	return i.store[creditCardNumber]
}
