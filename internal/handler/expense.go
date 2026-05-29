package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MrSafaran/expense-tracker/internal/model"
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

func (h *ExpenseHandler) ExpensesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		h.GetExpensesHandler(w, r)

	case http.MethodPost:
		h.CreateExpenseHandler(w, r)

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
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

func (h *ExpenseHandler) CreateExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var req model.CreateExpenseRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid request body",
			http.StatusBadRequest,
		)

		return
	}

	expense, err := h.service.CreateExpense(req)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(expense)
}

func (h *ExpenseHandler) ExpenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		h.GetExpenseByIDHandler(w, r)

	case http.MethodDelete:
		h.DeleteExpenseHandler(w, r)

	default:
		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

func (h *ExpenseHandler) DeleteExpenseHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(
		r.URL.Path,
		"/expenses/",
	)

	id, err := strconv.Atoi(path)

	if err != nil {
		http.Error(
			w,
			"invalid expense id",
			http.StatusBadRequest,
		)
		return
	}

	err = h.service.DeleteExpense(id)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ExpenseHandler) GetExpenseByIDHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(
		r.URL.Path,
		"/expenses/",
	)

	id, err := strconv.Atoi(path)

	if err != nil {
		http.Error(
			w,
			"invalid expense id",
			http.StatusBadRequest,
		)
		return
	}

	expense, err := h.service.GetExpenseByID(id)

	if err != nil {
		http.Error(
			w,
			"expense not found",
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(expense)
}
