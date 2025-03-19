package repositories

import (
	"books-app/internal/domain"
	"books-app/pkg"
	"errors"
)

type BookMemoryRepository struct {
	books     map[uint]domain.Book
	idCounter uint
}

func NewBookMemoryRepository() *BookMemoryRepository {
	return &BookMemoryRepository{
		books:     make(map[uint]domain.Book),
		idCounter: 1,
	}
}

func (r *BookMemoryRepository) CreateBook(book *domain.Book) (*domain.Book, error) {

	book.ID = r.idCounter
	r.books[r.idCounter] = *book
	r.idCounter++

	return book, nil

}

func (r *BookMemoryRepository) ReadBooks() ([]*domain.Book, error) {

	books := []*domain.Book{}
	for _, book := range r.books {
		books = append(books, &book)
	}

	return books, nil

}

func (r *BookMemoryRepository) ReadBook(id uint) (*domain.Book, error) {

	book, exists := r.books[id]
	if !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	return &book, nil

}

func (r *BookMemoryRepository) UpdateBook(id uint, book *domain.Book) (*domain.Book, error) {

	if _, exists := r.books[id]; !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	book.ID = id
	r.books[id] = *book

	return book, nil

}

func (r *BookMemoryRepository) DeleteBook(id uint) error {

	if _, exists := r.books[id]; !exists {
		return errors.New(pkg.ErrNotFound)
	}

	delete(r.books, id)

	return nil

}
