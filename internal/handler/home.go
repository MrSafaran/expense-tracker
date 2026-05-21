package handler

import (
	"encoding/json"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := HomeResponse{
		Message: "Welcome to Expense Tracker API",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}