package main

import (
	"fmt"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/handler"
)

func main() {
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/expenses", handler.ExpensesHandler)
	// http.HandleFunc("/expenses", handler.GetExpensesHandler)
	// http.HandleFunc("/expenses", handler.CreateExpenseHandler)

	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}