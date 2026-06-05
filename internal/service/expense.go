package service

import (
	"errors"

	"github.com/MrSafaran/expense-tracker/internal/model"

)

type ExpenseService struct {
	repo ExpenseRepository
}

func NewExpenseService(repo ExpenseRepository) *ExpenseService {
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

func (s *ExpenseService) GetExpenseByID(id int) (model.Expense, error) {

	return s.repo.GetExpenseByID(id)
}

func (s *ExpenseService) DeleteExpense(id int) error {

	return s.repo.DeleteExpense(id)
}

func (s *ExpenseService) UpdateExpense(id int,req model.CreateExpenseRequest) (model.Expense, error) {
	if req.Title == "" {
		return model.Expense{}, errors.New("title is required")
	}

	if req.Amount <= 0 {
		return model.Expense{}, errors.New("amount must be positive")
	}

	expense := model.Expense{
		ID:       id,
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}

	return s.repo.UpdateExpense(expense)
}