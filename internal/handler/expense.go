package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/MrSafaran/expense-tracker/internal/service"
)

func ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		GetExpensesHandler(w, r)

	case http.MethodPost:
		CreateExpenseHandler(w, r)

	case http.MethodDelete:
		DeleteExpenseHandler(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func ExpenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		GetExpenseByIDHandler(w, r)

	case http.MethodDelete:
		DeleteExpenseHandler(w, r)
		
	case http.MethodPut:
		UpdateExpenseHandler(w, r)

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

func GetExpenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	
	idStr := strings.TrimPrefix(path, "/expenses/")
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	expense, found := service.GetExpenseByID(id)
	if !found {
		http.Error(w, "expense not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(expense)
}

func DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	idStr := strings.TrimPrefix(path, "/expenses/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	deleted := service.DeleteExpense(id)
	if !deleted {
		http.Error(w, "expense not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func UpdateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	idStr := strings.TrimPrefix(path, "/expenses/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid expense id", http.StatusBadRequest)
		return
	}

	var req model.CreateExpenseRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	expense, err := service.UpdateExpense(id, req)
	if err != nil {

		if err.Error() == "expense not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(expense)
}
