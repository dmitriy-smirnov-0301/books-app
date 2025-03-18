package handlers

import (
	"books-app/internal/domain"
	"books-app/pkg"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BookHandlers struct {
	Repo domain.BookRepository
}

func NewBookHandlers(repo domain.BookRepository) *BookHandlers {
	return &BookHandlers{Repo: repo}
}

func (bh *BookHandlers) CreateBook(ctx echo.Context) error {

	var book domain.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidRequest])
	}

	newBook, err := bh.Repo.CreateBook(&book)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, newBook)

}

func (bh *BookHandlers) ReadBooks(ctx echo.Context) error {

	books, err := bh.Repo.ReadBooks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, books)

}

func (bh *BookHandlers) ReadBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	book, err := bh.Repo.ReadBook(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusOK, book)

}

func (bh *BookHandlers) UpdateBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	var book domain.Book
	if err := ctx.Bind(&book); err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidRequest])
	}

	updatedBook, err := bh.Repo.UpdateBook(id, &book)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusOK, updatedBook)

}

func (bh *BookHandlers) DeleteBook(ctx echo.Context) error {

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, pkg.Response[pkg.ErrInvalidID])
	}

	if err := bh.Repo.DeleteBook(id); err != nil {
		return ctx.JSON(http.StatusNotFound, pkg.Response[pkg.ErrNotFound])
	}

	return ctx.JSON(http.StatusNoContent, nil)

}
