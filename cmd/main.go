package main

import (
	"books-app/internal/controllers/handlers"
	"books-app/internal/controllers/router"
	"books-app/internal/repositories"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	repository := repositories.NewMemoryStorage()
	handlers := handlers.NewBookHandlers(repository)

	router.RegisterRoutes(e, handlers)

	e.Logger.Fatal(e.Start(":8080"))

}
