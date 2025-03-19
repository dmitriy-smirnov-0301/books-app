package main

import (
	_ "books-app/docs"
	"books-app/internal/controllers/handlers"
	"books-app/internal/controllers/router"
	"books-app/internal/db"
	"books-app/internal/repositories"
	"log"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Books API
// @version 1.0
// @description API для управления библиотекой книг
// @host localhost:8080
// @BasePath /
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))

}
