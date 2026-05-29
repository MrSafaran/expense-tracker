package router

import (
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/handler"
	"github.com/MrSafaran/expense-tracker/internal/middleware"
)

func RegisterRoutes(expenseHandler *handler.ExpenseHandler) {

	http.HandleFunc(
		"/",
		middleware.RecoveryMiddleware(
			middleware.LoggingMiddleware(
				handler.HomeHandler,
			),
		),
	)

	http.HandleFunc(
		"/expenses",
		middleware.RecoveryMiddleware(
			middleware.LoggingMiddleware(
				expenseHandler.ExpensesHandler,
			),
		),
	)
}
