package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreditCardValidator(t *testing.T) {
	t.Run("validating valid numbers for Luhn algorithm", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/credit_card_number/3379 5135 6110 8795", nil)
		response := httptest.NewRecorder()

		CreditCardValidatorServer(response, request)

		got := response.Body.String()
		want := "true"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("catching invalid numbers for Luhn algorithm", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/credit_card_number/3379 5135 6110 8794", nil)
		response := httptest.NewRecorder()

		CreditCardValidatorServer(response, request)

		got := response.Body.String()
		want := "false"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
