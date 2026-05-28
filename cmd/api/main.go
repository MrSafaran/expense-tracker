package main

import (
	"fmt"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/config"
	"github.com/MrSafaran/expense-tracker/internal/database"
	"github.com/MrSafaran/expense-tracker/internal/handler"
	"github.com/MrSafaran/expense-tracker/internal/repository"
	"github.com/MrSafaran/expense-tracker/internal/router"
	"github.com/MrSafaran/expense-tracker/internal/service"
)

func main() {

	cfg := config.Load()

	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	expenseRepo := repository.NewExpenseRepository(db)

	expenseService := service.NewExpenseService(expenseRepo)

	expenseHandler := handler.NewExpenseHandler(expenseService)

	router.RegisterRoutes(expenseHandler)

	fmt.Println("Server is running on :" + cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		panic(err)
	}
}