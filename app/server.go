package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"unicode"
)

type CreditCardStore interface {
	GetCardValidation(number string) bool
}

type CreditCardValidatorServer struct {
	store CreditCardStore
}

func (c *CreditCardValidatorServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	creditCardPayload, _ := io.ReadAll(r.Body)
	m := make(map[string]string)
	json.Unmarshal(creditCardPayload, &m)

	isValid := c.store.GetCardValidation(m["CreditCardNumber"])

	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
	}

	fmt.Fprint(w, isValid)
}

func GetCardValidation(number string) bool {
	var factor int = 2
	var sum int = 0
	var product int = 1
	for _, ch := range number {
		if unicode.IsDigit(ch) {
			product = (int(ch) - '0') * factor

			if product >= 10 {
				sum = sum + 1 + product%10
			} else {
				sum = sum + product
			}

			if factor == 2 {
				factor = 1
			} else {
				factor = 2
			}
		}
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
