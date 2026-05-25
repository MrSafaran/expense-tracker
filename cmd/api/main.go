package main

import (
	"fmt"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/router"
)

func main() {
	router.RegisterRoutes()

	fmt.Println("Server is running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}