package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/MrSafaran/expense-tracker/internal/model"
)

type FakeExpenseService struct {
	LastRequest         model.CreateExpenseRequest
	CreateExpenseError  error
	LastID              int
	GetExpenseByIDError error
	DeleteExpenseError  error
	UpdateExpenseError  error
	GetExpensesCalled   bool
	CreateExpenseCalled bool
}

// CreateExpense implements [ExpenseService].
func (f *FakeExpenseService) CreateExpense(req model.CreateExpenseRequest) (model.Expense, error) {
	f.LastRequest = req
	f.CreateExpenseCalled = true

	if f.CreateExpenseError != nil {
		return model.Expense{}, f.CreateExpenseError
	}

	return model.Expense{
		ID:       1,
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}, nil
}

// DeleteExpense implements [ExpenseService].
func (f *FakeExpenseService) DeleteExpense(id int) error {
	f.LastID = id

	if f.DeleteExpenseError != nil {
		return f.DeleteExpenseError
	}

	return nil
}

// GetExpenseByID implements [ExpenseService].
func (f *FakeExpenseService) GetExpenseByID(id int) (model.Expense, error) {
	f.LastID = id

	if f.GetExpenseByIDError != nil {
		return model.Expense{}, f.GetExpenseByIDError
	}

	return model.Expense{
		ID:       id,
		Title:    "Coffee",
		Amount:   10,
		Category: "Food",
	}, nil
}

// UpdateExpense implements [ExpenseService].
func (f *FakeExpenseService) UpdateExpense(id int, req model.CreateExpenseRequest) (model.Expense, error) {
	f.LastID = id
	f.LastRequest = req

	if f.UpdateExpenseError != nil {
		return model.Expense{}, f.UpdateExpenseError
	}

	return model.Expense{
		ID:       id,
		Title:    req.Title,
		Amount:   req.Amount,
		Category: req.Category,
	}, nil
}

func (f *FakeExpenseService) GetExpenses() ([]model.Expense, error) {
	f.GetExpensesCalled = true

	return []model.Expense{
		{
			ID:       1,
			Title:    "Coffee",
			Amount:   10,
			Category: "Food",
		},
	}, nil
}

func TestGetExpensesHandler(t *testing.T) {

	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.GetExpensesHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response []model.Expense
	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 expense, got %d",
			len(response),
		)
	}

	if response[0].ID != 1 {
		t.Fatalf(
			"expected id 1, got %d",
			response[0].ID,
		)
	}

	if response[0].Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			response[0].Title,
		)
	}

	if response[0].Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			response[0].Amount,
		)
	}

	if response[0].Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			response[0].Category,
		)
	}

}

func TestCreateExpenseHandler(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/expenses",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.CreateExpenseHandler(rr, req)

	if fakeService.LastRequest.Title != "Coffee" {
		t.Fatal("wrong title passed to service")
	}

	if fakeService.LastRequest.Amount != 10 {
		t.Fatal("wrong amount passed to service")
	}

	if fakeService.LastRequest.Category != "Food" {
		t.Fatal("wrong category passed to service")
	}

	if rr.Code != http.StatusCreated {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response model.Expense

	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if response.ID != 1 {
		t.Fatalf(
			"expected id 1, got %d",
			response.ID,
		)
	}

	if response.Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			response.Title,
		)
	}

	if response.Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			response.Amount,
		)
	}

	if response.Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			response.Category,
		)
	}

}

func TestCreateExpenseHandler_InvalidJSON(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
	"title":
	`
	req := httptest.NewRequest(
		http.MethodPost,
		"/expenses",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.CreateExpenseHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "invalid request body") {
		t.Fatal("wrong error message")
	}
}

func TestCreateExpenseHandler_ServiceError(t *testing.T) {
	fakeService := &FakeExpenseService{
		CreateExpenseError: errors.New("title is required"),
	}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/expenses",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.CreateExpenseHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "title is required") {
		t.Fatal("wrong error message")
	}
}

func TestGetExpenseByIDHandler_Success(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.GetExpenseByIDHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response model.Expense

	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if response.ID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			response.ID,
		)
	}

	if response.Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			response.Title,
		)
	}

	if response.Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			response.Amount,
		)
	}

	if response.Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			response.Category,
		)
	}
}

func TestGetExpenseByIDHandler_InvalidID(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses/abc",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.GetExpenseByIDHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "invalid expense id") {
		t.Fatal("wrong error message")
	}
}

func TestGetExpenseByIDHandler_NotFound(t *testing.T) {
	fakeService := &FakeExpenseService{
		GetExpenseByIDError: errors.New("not found"),
	}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.GetExpenseByIDHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "expense not found") {
		t.Fatal("wrong error message")
	}
}

func TestDeleteExpenseHandler_Success(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.DeleteExpenseHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}

	if rr.Code != http.StatusNoContent {
		t.Fatalf(
			"expected status 204, got %d",
			rr.Code,
		)
	}

	if rr.Body.String() != "" {
		t.Fatal("expected empty response body")
	}
}

func TestDeleteExpenseHandler_InvalidID(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/expenses/abc",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.DeleteExpenseHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "invalid expense id") {
		t.Fatal("wrong error message")
	}
}

func TestDeleteExpenseHandler_ServiceError(t *testing.T) {
	fakeService := &FakeExpenseService{
		DeleteExpenseError: errors.New("expense not found"),
	}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.DeleteExpenseHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}

	if rr.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status 404, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "expense not found") {
		t.Fatal("wrong error message")
	}
}

func TestUpdateExpenseHandler_Success(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/expenses/123",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.UpdateExpenseHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}

	if fakeService.LastRequest.Title != "Coffee" {
		t.Fatal("wrong title passed to service")
	}

	if fakeService.LastRequest.Amount != 10 {
		t.Fatal("wrong amount passed to service")
	}

	if fakeService.LastRequest.Category != "Food" {
		t.Fatal("wrong category passed to service")
	}

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rr.Code,
		)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Fatal("wrong content type")
	}

	var response model.Expense

	err := json.NewDecoder(rr.Body).Decode(&response)

	if err != nil {
		t.Fatalf(
			"failed to decode response: %v",
			err,
		)
	}

	if response.ID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			response.ID,
		)
	}

	if response.Title != "Coffee" {
		t.Fatalf(
			"expected title Coffee, got %s",
			response.Title,
		)
	}

	if response.Amount != 10 {
		t.Fatalf(
			"expected amount 10, got %f",
			response.Amount,
		)
	}

	if response.Category != "Food" {
		t.Fatalf(
			"expected category Food, got %s",
			response.Category,
		)
	}
}

func TestUpdateExpenseHandler_InvalidID(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodPut,
		"/expenses/abc",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.UpdateExpenseHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "invalid expense id") {
		t.Fatal("wrong error message")
	}
}

func TestUpdateExpenseHandler_InvalidJSON(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/expenses/123",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.UpdateExpenseHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "invalid request body") {
		t.Fatal("wrong error message")
	}
}

func TestUpdateExpenseHandler_ServiceError(t *testing.T) {
	fakeService := &FakeExpenseService{
		UpdateExpenseError: errors.New("expense not found"),
	}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/expenses/123",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.UpdateExpenseHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}

	if rr.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "expense not found") {
		t.Fatal("wrong error message")
	}
}

func TestExpensesHandler_MethodNotAllowed(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodPatch,
		"/expenses",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.ExpensesHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status 405, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "method not allowed") {
		t.Fatal("wrong error message")
	}
}

func TestExpensesHandler_GET(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.ExpensesHandler(rr, req)

	if !fakeService.GetExpensesCalled {
		t.Fatal("GetExpenses was not called")
	}
}

func TestExpensesHandler_POST(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/expenses",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.ExpensesHandler(rr, req)

	if !fakeService.CreateExpenseCalled {
		t.Fatal("CreateExpenses was not called")
	}
}

func TestExpenseByIDHandler_MethodNotAllowed (t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodPatch,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.ExpenseByIDHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status 405, got %d",
			rr.Code,
		)
	}

	if !strings.Contains(rr.Body.String(), "method not allowed") {
		t.Fatal("wrong error message")
	}
}

func TestExpenseByIDHandler_GET(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodGet,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.ExpenseByIDHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}
}

func TestExpenseByIDHandler_DELETE(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/expenses/123",
		nil,
	)

	rr := httptest.NewRecorder()

	handler.ExpenseByIDHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}
}

func TestExpenseByIDHandler_PUT(t *testing.T) {
	fakeService := &FakeExpenseService{}

	handler := NewExpenseHandler(fakeService)

	body := `{
		"title":"Coffee",
		"amount":10,
		"category":"Food"
	}`

	req := httptest.NewRequest(
		http.MethodPut,
		"/expenses/123",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	handler.ExpenseByIDHandler(rr, req)

	if fakeService.LastID != 123 {
		t.Fatalf(
			"expected id 123, got %d",
			fakeService.LastID,
		)
	}
}