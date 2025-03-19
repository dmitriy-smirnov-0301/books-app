package handlers

import (
	"books-app/internal/domain"
	"books-app/pkg"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandlers struct {
	repo domain.BookRepository
}

func NewBookHandlers(repo domain.BookRepository) *BookHandlers {
	return &BookHandlers{
		repo: repo,
	}
}

// CreateBook добавляет новую книгу
// @Summary Добавить книгу
// @Description Добавляет новую книгу в базу данных
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body domain.Book true "Данные книги"
// @Success 201 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid request payload"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /books [post]
func (h *BookHandlers) CreateBook(ctx echo.Context) error {

	var book domain.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidRequest])
	}

	newBook, err := h.repo.CreateBook(&book)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, newBook)

}

// ReadBooks возвращает список всех книг
// @Summary Получить список книг
// @Description Возвращает все книги из базы данных
// @Tags books
// @Produce json
// @Success 200 {array} domain.Book
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /books [get]
func (h *BookHandlers) ReadBooks(ctx echo.Context) error {

	books, err := h.repo.ReadBooks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, books)

}

// ReadBook возвращает книгу по ID
// @Summary Получить книгу
// @Description Возвращает книгу по её ID
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid book ID"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [get]
func (h *BookHandlers) ReadBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	book, err := h.repo.ReadBook(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusOK, book)

}

// UpdateBook обновляет книгу по ID
// @Summary Обновить книгу
// @Description Обновляет информацию о книге в базе данных по её ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Param book body domain.Book true "Обновлённые данные книги"
// @Success 200 {object} domain.Book
// @Failure 400 {object} map[string]string "Invalid book ID or request payload"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [put]
func (h *BookHandlers) UpdateBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	var book domain.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidRequest])
	}

	updatedBook, err := h.repo.UpdateBook(uint(id), &book)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusOK, updatedBook)

}

// DeleteBook удаляет книгу по ID
// @Summary Удалить книгу
// @Description Удаляет книгу из базы данных по её ID
// @Tags books
// @Produce json
// @Param id path int true "ID книги"
// @Success 204 "Book deleted successfully"
// @Failure 400 {object} map[string]string "Invalid book ID"
// @Failure 404 {object} map[string]string "Book not found"
// @Router /books/{id} [delete]
func (h *BookHandlers) DeleteBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	if err := h.repo.DeleteBook(uint(id)); err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusNoContent, nil)

}
