package handlers

import (
	"books-app/internal/domain"
	"books-app/internal/repositories"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setupTestHandlers() *BookHandlers {
	repo := repositories.NewMockBookRepository()
	return NewBookHandlers(repo)
}

func TestCreateBook(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	book := domain.Book{
		Title:           "Test Book",
		Author:          "Author Test",
		PublicationYear: 2023,
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := h.CreateBook(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var createdBook domain.Book
	err = json.Unmarshal(rec.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, createdBook.Title)
}

func TestReadBooks(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := h.ReadBooks(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestReadBook_NotFound(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	req := httptest.NewRequest(http.MethodGet, "/books/999", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("999")

	err := h.ReadBook(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	var response map[string]string
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Book not found", response["error"])
}

func TestUpdateBook_Success(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	book := domain.Book{
		Title:           "Test Book",
		Author:          "Test Author",
		PublicationYear: 2023,
	}
	h.repo.CreateBook(&book)

	updateData := domain.Book{
		Title: "Updated Title",
	}
	body, _ := json.Marshal(updateData)
	req := httptest.NewRequest(http.MethodPut, "/books/"+strconv.Itoa(int(book.ID)), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(book.ID)))

	err := h.UpdateBook(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var updatedBook domain.Book
	err = json.Unmarshal(rec.Body.Bytes(), &updatedBook)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updatedBook.Title)
}

func TestDeleteBook_Success(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	book := domain.Book{
		Title:           "Test Book",
		Author:          "Test Author",
		PublicationYear: 2023,
	}
	h.repo.CreateBook(&book)

	req := httptest.NewRequest(http.MethodDelete, "/books/"+strconv.Itoa(int(book.ID)), nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(book.ID)))

	err := h.DeleteBook(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestDeleteBook_NotFound(t *testing.T) {
	e := echo.New()
	h := setupTestHandlers()

	req := httptest.NewRequest(http.MethodDelete, "/books/999", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("999")

	err := h.DeleteBook(ctx)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}
