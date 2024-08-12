package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreditCardValidator(t *testing.T) {
	// 3379 5135 6110 8795
	// 2769 1483 0405 9987
	t.Run("validating valid numbers for Luhn algorithm", func(t *testing.T) {
		request := newGetValidationRequest("3379 5135 6110 8795")
		response := httptest.NewRecorder()

		CreditCardValidatorServer(response, request)

		assertResponseBody(t, response.Body.String(), "true")
	})
	// 3379 5135 6110 8794
	// 2769 1483 0405 9986
	t.Run("catching invalid numbers for Luhn algorithm", func(t *testing.T) {
		request := newGetValidationRequest("3379 5135 6110 8794")
		response := httptest.NewRecorder()

		CreditCardValidatorServer(response, request)

		assertResponseBody(t, response.Body.String(), "false")
	})
}

func newGetValidationRequest(creditCardNumber string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/credit_card_number/%s", creditCardNumber), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
