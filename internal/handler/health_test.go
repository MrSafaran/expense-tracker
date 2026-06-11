package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	rr := httptest.NewRecorder()

	HealthHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response map[string]string
	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if response["status"] != "ok" {
		t.Fatalf(
			"expected status ok, got %s",
			response["status"],
		)
	}
}
