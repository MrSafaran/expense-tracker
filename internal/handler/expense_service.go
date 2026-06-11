package handler

import "github.com/MrSafaran/expense-tracker/internal/model"

type ExpenseService interface {
	GetExpenses() ([]model.Expense, error)
	GetExpenseByID(id int) (model.Expense, error)
	CreateExpense(req model.CreateExpenseRequest) (model.Expense, error)
	UpdateExpense(id int, req model.CreateExpenseRequest) (model.Expense, error)
	DeleteExpense(id int) error
}
