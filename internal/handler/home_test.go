package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	rr := httptest.NewRecorder()

	HomeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response HomeResponse
	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if response.Message != "Welcome to Expense Tracker API" {
		t.Fatalf(
			"expected Welcome message, but got %s",
			response.Message,
		)
	}
}
