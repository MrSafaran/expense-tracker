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

func GetExpenseByID(id int) (model.Expense, bool) {
	for _, expense := range expenses {
		if expense.ID == id {
			return expense, true
		}
	}

	return model.Expense{}, false
}

func DeleteExpense(id int) bool {
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return true
		}
	}

	return false
}

func UpdateExpense(updatedExpense model.Expense) bool {
	for i, expense := range expenses {
		if expense.ID == updatedExpense.ID {
			expenses[i] = updatedExpense
			return true
		}
	}

	return false
}