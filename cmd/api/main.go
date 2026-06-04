package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr: ":" + cfg.Port,
	}

	go func() {
		fmt.Println("Server is running on :" + cfg.Port)

		err := server.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-quit

	fmt.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("Server stopped gracefully")
}
