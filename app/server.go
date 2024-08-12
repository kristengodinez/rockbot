package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

func CreditCardValidatorServer(w http.ResponseWriter, r *http.Request) {
	creditCardNumber := strings.TrimPrefix(r.URL.Path, "/credit_card_number/")

	var factor int = 2
	var sum int = 0
	var product int = 1
	for _, ch := range creditCardNumber {
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
		fmt.Fprint(w, "true")
		return
	}

	fmt.Fprint(w, "false")

}
