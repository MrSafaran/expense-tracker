package repository

import (
	"context"
	"errors"

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

func (r *ExpenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	query := `
		INSERT INTO expenses (
			title,
			amount,
			category
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		expense.Title,
		expense.Amount,
		expense.Category,
	).Scan(&expense.ID)

	if err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (r *ExpenseRepository) GetExpenseByID(id int) (model.Expense, error) {
	var expense model.Expense

	query := `
		SELECT
			id,
			title,
			amount,
			category
		FROM expenses
		WHERE id = $1
	`

	err := r.db.QueryRow(
		context.Background(),
		query,
		id,
	).Scan(
		&expense.ID,
		&expense.Title,
		&expense.Amount,
		&expense.Category,
	)

	if err != nil {
		return model.Expense{}, err
	}

	return expense, nil
}

func (r *ExpenseRepository) DeleteExpense(id int) error {
	result, err := r.db.Exec(
		context.Background(),
		"DELETE FROM expenses WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("expense not found")
	}

	return nil
}
