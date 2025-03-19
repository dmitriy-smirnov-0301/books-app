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

func (h *BookHandlers) ReadBooks(ctx echo.Context) error {

	books, err := h.repo.ReadBooks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, books)

}

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
