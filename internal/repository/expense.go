package repository

import "github.com/MrSafaran/expense-tracker/internal/model"

var expenses = []model.Expense{
	{
		ID:       1,
		Title:    "Groceries",
		Amount:   120.50,
		Category: "Food",
	},
	{
		ID:       2,
		Title:    "Internet",
		Amount:   45.00,
		Category: "Bills",
	},
}

func GetExpenses() []model.Expense {
	return expenses
}

func SaveExpense(expense model.Expense) {
	expenses = append(expenses, expense)
}
