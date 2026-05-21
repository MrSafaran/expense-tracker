package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/MrSafaran/expense-tracker/internal/service"
)

var expenses = []model.Expense{
	{
		ID:       1,
		Title:    "Groceries",
		Amount:   120.50,
		Category: "Food",
	},
	{
		ID:       2,
		Title:    "Internet",
		Amount:   45.00,
		Category: "Bills",
	},
}

func ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		GetExpensesHandler(w, r)

	case http.MethodPost:
		CreateExpenseHandler(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	expenses := service.GetExpenses()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(expenses)
}

func CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateExpenseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	expense, err := service.CreateExpense(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(expense)
}
