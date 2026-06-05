package service

import "github.com/MrSafaran/expense-tracker/internal/model"

type ExpenseRepository interface {
	GetExpenses() ([]model.Expense, error)
	CreateExpense(expense model.Expense) (model.Expense, error)
	GetExpenseByID(id int) (model.Expense, error)
	DeleteExpense(id int) error
	UpdateExpense(expense model.Expense) (model.Expense, error)
}