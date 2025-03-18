package repositories

import (
	"books-app/internal/domain"
	"books-app/pkg"
	"errors"
)

type MemoryStorage struct {
	Books     map[int]domain.Book
	IDCounter int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Books:     make(map[int]domain.Book),
		IDCounter: 1,
	}
}

func (ms *MemoryStorage) CreateBook(book *domain.Book) (*domain.Book, error) {

	book.ID = ms.IDCounter
	ms.Books[ms.IDCounter] = *book
	ms.IDCounter++

	return book, nil

}

func (ms *MemoryStorage) ReadBooks() ([]*domain.Book, error) {

	bookList := []*domain.Book{}
	for _, book := range ms.Books {
		bookList = append(bookList, &book)
	}

	return bookList, nil

}

func (ms *MemoryStorage) ReadBook(id int) (*domain.Book, error) {

	book, exists := ms.Books[id]
	if !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	return &book, nil

}

func (ms *MemoryStorage) UpdateBook(id int, book *domain.Book) (*domain.Book, error) {

	if _, exists := ms.Books[id]; !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	book.ID = id
	ms.Books[id] = *book

	return book, nil

}

func (ms *MemoryStorage) DeleteBook(id int) error {

	if _, exists := ms.Books[id]; !exists {
		return errors.New(pkg.ErrNotFound)
	}

	delete(ms.Books, id)

	return nil

}
