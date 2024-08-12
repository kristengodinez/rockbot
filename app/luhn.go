package main

import (
	"unicode"
)

func Luhn(number string) bool {
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
