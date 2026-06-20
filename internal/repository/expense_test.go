package repository

import (
	"context"
	"os"
	"testing"

	"github.com/MrSafaran/expense-tracker/internal/database"
	"github.com/MrSafaran/expense-tracker/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	err := godotenv.Load("../../.env")

	if err != nil {
		t.Fatalf(
			"failed to load env: %v",
			err,
		)
	}

	db, err := database.NewConnection(
		os.Getenv("DATABASE_URL"),
	)

	if err != nil {
		t.Fatalf(
			"failed to connect database: %v",
			err,
		)
	}

	return db
}

func TestCreateExpense(t *testing.T) {

	db := setupTestDB(t)
	defer db.Close()

	repo := NewExpenseRepository(db)

	createdExpense, err := repo.CreateExpense(
		model.Expense{
			Title:    "Coffee",
			Amount:   10,
			Category: "Food",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	defer func() {
		_, err := db.Exec(
			context.Background(),
			"DELETE FROM expenses WHERE id = $1",
			createdExpense.ID,
		)

		if err != nil {
			t.Fatalf(
				"failed cleanup: %v",
				err,
			)
		}

	}()

	if createdExpense.ID == 0 {
		t.Fatal("expected non-zero id")
	}

	if createdExpense.CreatedAt.IsZero() {
		t.Fatal("expected created_at to be set")
	}

	if createdExpense.Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			createdExpense.Title,
		)
	}

	if createdExpense.Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			createdExpense.Amount,
		)
	}

	if createdExpense.Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			createdExpense.Category,
		)
	}

	fetchedExpense, err := repo.GetExpenseByID(
		createdExpense.ID,
	)

	if err != nil {
		t.Fatalf(
			"failed to get expense: %v",
			err,
		)
	}

	if fetchedExpense.ID != createdExpense.ID {
		t.Fatalf(
			"expected id %d, got %d",
			createdExpense.ID,
			fetchedExpense.ID,
		)
	}

	if fetchedExpense.Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			fetchedExpense.Title,
		)
	}

	if fetchedExpense.Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			fetchedExpense.Amount,
		)
	}

	if fetchedExpense.Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			fetchedExpense.Category,
		)
	}

	if fetchedExpense.CreatedAt.IsZero() {
		t.Fatal("expected created_at to be set")
	}

}

func TestUpdateExpense(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewExpenseRepository(db)

	createdExpense, err := repo.CreateExpense(
		model.Expense{
			Title:    "Coffee",
			Amount:   10,
			Category: "Food",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	defer func() {
		_, err := db.Exec(
			context.Background(),
			"DELETE FROM expenses WHERE id = $1",
			createdExpense.ID,
		)

		if err != nil {
			t.Fatalf(
				"failed cleanup: %v",
				err,
			)
		}

	}()

	updateExpense := createdExpense

	updateExpense.Title = "Tea"
	updateExpense.Amount = 20
	updateExpense.Category = "Drink"

	updatedExpense, err := repo.UpdateExpense(
		updateExpense,
	)

	if err != nil {
		t.Fatalf(
			"failed to update expense: %v",
			err,
		)
	}

	if updatedExpense.ID != createdExpense.ID {
		t.Fatalf(
			"expected id %d, got %d",
			createdExpense.ID,
			updatedExpense.ID,
		)
	}

	if updatedExpense.Title != "Tea" {
		t.Fatalf(
			"expected title Tea, got %s",
			updatedExpense.Title,
		)
	}

	if updatedExpense.Amount != 20 {
		t.Fatalf(
			"expected amount 20, got %f",
			updatedExpense.Amount,
		)
	}

	if updatedExpense.Category != "Drink" {
		t.Fatalf(
			"expected category Drink, got %s",
			updatedExpense.Category,
		)
	}

	fetchedExpense, err := repo.GetExpenseByID(
		createdExpense.ID,
	)

	if err != nil {
		t.Fatalf(
			"failed to get expense: %v",
			err,
		)
	}

	if fetchedExpense.ID != createdExpense.ID {
		t.Fatalf(
			"expected id %d, got %d",
			createdExpense.ID,
			fetchedExpense.ID,
		)
	}

	if fetchedExpense.Title != "Tea" {
		t.Fatalf(
			"expected title Tea, got %s",
			fetchedExpense.Title,
		)
	}

	if fetchedExpense.Amount != 20 {
		t.Fatalf(
			"expected amount 20, got %f",
			fetchedExpense.Amount,
		)
	}

	if fetchedExpense.Category != "Drink" {
		t.Fatalf(
			"expected category Drink, got %s",
			fetchedExpense.Category,
		)
	}

	if fetchedExpense.CreatedAt.IsZero() {
		t.Fatal("expected created_at to be set")
	}

}

func TestDeleteExpense(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewExpenseRepository(db)

	createdExpense, err := repo.CreateExpense(
		model.Expense{
			Title:    "Coffee",
			Amount:   10,
			Category: "Food",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	err = repo.DeleteExpense(
		createdExpense.ID,
	)

	if err != nil {
		t.Fatalf(
			"failed to Delete expense: %v",
			err,
		)
	}

	_, err = repo.GetExpenseByID(
		createdExpense.ID,
	)

	if err == nil {
		t.Fatalf(
			"expected expense to be deleted: %v",
			err,
		)
	}
}

func TestDeleteExpense_NotFound(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewExpenseRepository(db)

	err := repo.DeleteExpense(999999)

	if err == nil {
		t.Fatal(
			"expected error, got nil",
		)
	}

	if err.Error() != "expense not found" {
		t.Fatalf(
			"expected expense not found, got %v",
			err,
		)
	}
}

func TestGetExpenses(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewExpenseRepository(db)

	createdExpense1, err := repo.CreateExpense(
		model.Expense{
			Title:    "Coffee",
			Amount:   10,
			Category: "Food",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	createdExpense2, err := repo.CreateExpense(
		model.Expense{
			Title:    "Pizza",
			Amount:   20,
			Category: "Food",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	createdExpense3, err := repo.CreateExpense(
		model.Expense{
			Title:    "Book",
			Amount:   30,
			Category: "Objects",
		},
	)

	if err != nil {
		t.Fatalf(
			"failed to create expense: %v",
			err,
		)
	}

	defer func() {

		_, err := db.Exec(
			context.Background(),
			"DELETE FROM expenses WHERE id = $1",
			createdExpense1.ID,
		)
		if err != nil {
			t.Fatalf(
				"failed to cleanup expense1: %v",
				err,
			)
		}

		_, err = db.Exec(
			context.Background(),
			"DELETE FROM expenses WHERE id = $1",
			createdExpense2.ID,
		)
		if err != nil {
			t.Fatalf(
				"failed to cleanup expense2: %v",
				err,
			)
		}

		_, err = db.Exec(
			context.Background(),
			"DELETE FROM expenses WHERE id = $1",
			createdExpense3.ID,
		)
		if err != nil {
			t.Fatalf(
				"failed to cleanup expense3: %v",
				err,
			)
		}

	}()

	expenses, err := repo.GetExpenses()

	if err != nil {
		t.Fatalf(
			"failed to get expenses: %v",
			err,
		)
	}

	foundCoffee := false
	foundPizza := false
	foundBook := false

	for _, expense := range expenses {

		if expense.Title == "Coffee" {
			foundCoffee = true
		}

		if expense.Title == "Pizza" {
			foundPizza = true
		}

		if expense.Title == "Book" {
			foundBook = true
		}
	}

	if !foundCoffee {
		t.Fatal("Coffee not found")
	}

	if !foundPizza {
		t.Fatal("Pizza not found")
	}

	if !foundBook {
		t.Fatal("Book not found")
	}
}
