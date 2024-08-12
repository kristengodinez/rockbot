package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryCreditCardStore()
	server := CreditCardValidatorServer{store}
	jsonPayload := []byte(`{"CreditCardNumber": "3379 5135 6110 8795"}`)

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetValidationRequest(jsonPayload))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "true")
}
