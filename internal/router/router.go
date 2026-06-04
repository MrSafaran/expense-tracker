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

	http.HandleFunc(
		"/expenses/",
		middleware.RecoveryMiddleware(
			middleware.LoggingMiddleware(
				expenseHandler.ExpenseByIDHandler,
			),
		),
	)

	http.HandleFunc(
		"/health",
		middleware.RecoveryMiddleware(
			middleware.LoggingMiddleware(
				handler.HealthHandler,
			),
		),
	)
}
