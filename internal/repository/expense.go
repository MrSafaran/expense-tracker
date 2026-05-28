package repository

import (
	"context"

	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExpenseRepository struct {
	db *pgxpool.Pool
}

func NewExpenseRepository(db *pgxpool.Pool) *ExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

func (r *ExpenseRepository) GetExpenses() ([]model.Expense, error) {

	rows, err := r.db.Query(
		context.Background(),
		"SELECT id, title, amount, category FROM expenses",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	expenses := []model.Expense{}

	for rows.Next() {

		var expense model.Expense

		err := rows.Scan(
			&expense.ID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
		)

		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}

