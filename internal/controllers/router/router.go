package router

import (
	"books-app/internal/controllers/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, bookHandlers *handlers.BookHandlers) {
	e.POST("/books", bookHandlers.CreateBook)
	e.GET("/books", bookHandlers.ReadBooks)
	e.GET("/books/:id", bookHandlers.ReadBook)
	e.PUT("/books/:id", bookHandlers.UpdateBook)
	e.DELETE("/books/:id", bookHandlers.DeleteBook)
}
