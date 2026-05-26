package main

import (
	"fmt"
	"net/http"

	"github.com/MrSafaran/expense-tracker/internal/router"
	"github.com/MrSafaran/expense-tracker/internal/config"
)

func main() {
	router.RegisterRoutes()
	cfg := config.Load()

	fmt.Println("Server is running on :" + cfg.Port)

	err := http.ListenAndServe(":" + cfg.Port, nil)
	if err != nil {
		panic(err)
	}
}