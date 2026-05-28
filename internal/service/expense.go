package service

import (
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
