package service

import (
	"testing"

	"github.com/MrSafaran/expense-tracker/internal/model"
)

type FakeExpenseRepository struct {
	CreateCalled      bool
	GetExpensesCalled bool
	LastID            int
	DeleteCalled      bool
}

func (f *FakeExpenseRepository) GetExpenses() ([]model.Expense, error) {
	f.GetExpensesCalled = true

	return []model.Expense{}, nil
}

func (f *FakeExpenseRepository) CreateExpense(expense model.Expense) (model.Expense, error) {
	f.CreateCalled = true
	expense.ID = 1

	return expense, nil
}

func (f *FakeExpenseRepository) GetExpenseByID(id int) (model.Expense, error) {

	f.LastID = id

	return model.Expense{}, nil
}

func (f *FakeExpenseRepository) DeleteExpense(id int) error {

	f.DeleteCalled = true
	f.LastID = id

	return nil
}

func (f *FakeExpenseRepository) UpdateExpense(expense model.Expense) (model.Expense, error) {

	return expense, nil
}

func TestCreateExpense(t *testing.T) {
	tests := []struct {
		name        string
		req         model.CreateExpenseRequest
		wantErr     bool
		expectedErr string
	}{
		{
			name: "empty title",
			req: model.CreateExpenseRequest{
				Title:    "",
				Amount:   10,
				Category: "Food",
			},
			wantErr:     true,
			expectedErr: "title is required",
		},
		{
			name: "invalid amount",
			req: model.CreateExpenseRequest{
				Title:    "Coffee",
				Amount:   -1,
				Category: "Food",
			},
			wantErr:     true,
			expectedErr: "amount must be positive",
		},
		{
			name: "success",
			req: model.CreateExpenseRequest{
				Title:    "Coffee",
				Amount:   10,
				Category: "Food",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			repo := &FakeExpenseRepository{}

			service := NewExpenseService(repo)

			_, err := service.CreateExpense(tt.req)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantErr {
				if err.Error() != tt.expectedErr {
					t.Fatalf("expected error %q, got %q", tt.expectedErr, err.Error())
				}
			}
		})
	}
}

func TestUpdateExpense(t *testing.T) {
	tests := []struct {
		name        string
		req         model.CreateExpenseRequest
		wantErr     bool
		expectedErr string
	}{
		{
			name: "empty title",
			req: model.CreateExpenseRequest{
				Title:    "",
				Amount:   10,
				Category: "Food",
			},
			wantErr:     true,
			expectedErr: "title is required",
		},
		{
			name: "invalid amount",
			req: model.CreateExpenseRequest{
				Title:    "Coffee",
				Amount:   -1,
				Category: "Food",
			},
			wantErr:     true,
			expectedErr: "amount must be positive",
		},
		{
			name: "success",
			req: model.CreateExpenseRequest{
				Title:    "Coffee",
				Amount:   10,
				Category: "Food",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			repo := &FakeExpenseRepository{}

			service := NewExpenseService(repo)

			_, err := service.UpdateExpense(123, tt.req)

			if tt.wantErr && err == nil {
				t.Fatal("expected error but got nil")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.wantErr {
				if err.Error() != tt.expectedErr {
					t.Fatalf("expected error %q, got %q", tt.expectedErr, err.Error())
				}
			}
		})
	}
}

func TestGetExpenses(t *testing.T) {
	repo := &FakeExpenseRepository{}

	service := NewExpenseService(repo)

	_, err := service.GetExpenses()

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if !repo.GetExpensesCalled {
		t.Fatal("GetExpenses was not called")
	}
}

func TestGetExpenseByID(t *testing.T) {
	repo := &FakeExpenseRepository{}

	service := NewExpenseService(repo)

	_, err := service.GetExpenseByID(123)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if repo.LastID != 123 {
		t.Fatal("GetExpenseByID was not called by correct ID")
	}
}

func TestDeleteExpense(t *testing.T) {

	repo := &FakeExpenseRepository{}

	service := NewExpenseService(repo)

	err := service.DeleteExpense(123)

	if err != nil {
		t.Fatalf(
			"unexpected error: %v",
			err,
		)
	}

	if !repo.DeleteCalled {
		t.Fatal(
			"DeleteExpense was not called",
		)
	}

	if repo.LastID != 123 {
		t.Fatal(
			"wrong id passed to repository",
		)
	}
}
