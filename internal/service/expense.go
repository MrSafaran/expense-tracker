package service

import (
	"errors"

	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/MrSafaran/expense-tracker/internal/repository"
)

func GetExpenses() []model.Expense {
	return repository.GetExpenses()
}

func CreateExpense(req model.CreateExpenseRequest) (model.Expense, error) {
	expenses := repository.GetExpenses()

	if req.Title == "" {
		return model.Expense{}, errors.New("title is required")
	}

	if req.Amount <= 0 {
		return model.Expense{}, errors.New("amount must be positive")
	}

	expense := model.Expense{
		ID:       len(expenses) + 1,
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}

	repository.SaveExpense(expense)

	return expense, nil
}
