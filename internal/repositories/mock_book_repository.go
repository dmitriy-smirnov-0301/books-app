package repositories

import (
	"books-app/internal/domain"
	"books-app/pkg"
	"errors"
	"sync"
)

type MockBookRepository struct {
	mu     sync.Mutex
	books  map[uint]*domain.Book
	nextID uint
}

func NewMockBookRepository() *MockBookRepository {
	return &MockBookRepository{
		books:  make(map[uint]*domain.Book),
		nextID: 1,
	}
}

func (m *MockBookRepository) CreateBook(book *domain.Book) (*domain.Book, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	book.ID = m.nextID
	m.books[m.nextID] = book
	m.nextID++

	return book, nil
}

func (m *MockBookRepository) ReadBooks() ([]*domain.Book, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var books []*domain.Book
	for _, book := range m.books {
		books = append(books, book)
	}

	return books, nil
}

func (m *MockBookRepository) ReadBook(id uint) (*domain.Book, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	book, exists := m.books[id]
	if !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	return book, nil
}

func (m *MockBookRepository) UpdateBook(id uint, book *domain.Book) (*domain.Book, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	existingBook, exists := m.books[id]
	if !exists {
		return nil, errors.New(pkg.ErrNotFound)
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.PublicationYear != 0 {
		existingBook.PublicationYear = book.PublicationYear
	}

	return existingBook, nil
}

func (m *MockBookRepository) DeleteBook(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.books[id]; !exists {
		return errors.New(pkg.ErrNotFound)
	}

	delete(m.books, id)
	return nil
}
