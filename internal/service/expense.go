package service

import (
	 "errors"
	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/MrSafaran/expense-tracker/internal/repository"
)

type ExpenseService struct {
	repo *repository.ExpenseRepository
}

func NewExpenseService(repo *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		repo: repo,
	}
}

func (s *ExpenseService) GetExpenses() ([]model.Expense, error) {
	return s.repo.GetExpenses()
}

func (s *ExpenseService) CreateExpense(req model.CreateExpenseRequest) (model.Expense, error) {
	if req.Title == "" {
		return model.Expense{},
			errors.New("title is required")
	}

	if req.Amount <= 0 {
		return model.Expense{},
			errors.New("amount must be positive")
	}

	expense := model.Expense{
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}

	return s.repo.CreateExpense(expense)
}