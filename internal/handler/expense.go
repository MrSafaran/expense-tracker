package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/service"
)

type ExpenseHandler struct {
	service *service.ExpenseService
}

func NewExpenseHandler(service *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		service: service,
	}
}

func (h *ExpenseHandler) GetExpensesHandler(w http.ResponseWriter, r *http.Request) {

	expenses, err := h.service.GetExpenses()

	if err != nil {
		http.Error(
			w,
			"failed to get expenses",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(expenses)
}