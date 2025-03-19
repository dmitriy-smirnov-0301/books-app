package main

import (
	"books-app/internal/controllers/handlers"
	"books-app/internal/controllers/router"
	"books-app/internal/db"
	"books-app/internal/repositories"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	database := db.NewDatabase()

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	connection := database.GetConnection()

	e := echo.New()

	// repository := repositories.NewBookMemoryRepository()
	repository := repositories.NewBookPostgresRepository(connection)
	handlers := handlers.NewBookHandlers(repository)

	router.RegisterRoutes(e, handlers)

	e.Logger.Fatal(e.Start(":8080"))

}
