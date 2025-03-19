package repositories

import (
	"books-app/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type BookPostgresRepository struct {
	db *gorm.DB
}

func NewBookPostgresRepository(db *gorm.DB) *BookPostgresRepository {
	return &BookPostgresRepository{
		db: db,
	}
}

func (r *BookPostgresRepository) CreateBook(book *domain.Book) (*domain.Book, error) {

	if err := r.db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil

}

func (r *BookPostgresRepository) ReadBooks() ([]*domain.Book, error) {

	books := []*domain.Book{}
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil

}

func (r *BookPostgresRepository) ReadBook(id uint) (*domain.Book, error) {

	book := &domain.Book{}
	if err := r.db.First(book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return book, nil

}

func (r *BookPostgresRepository) UpdateBook(id uint, book *domain.Book) (*domain.Book, error) {

	existingBook := &domain.Book{}
	if err := r.db.First(existingBook, id).Error; err != nil {
		return nil, errors.New("book not found")
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

	if err := r.db.Save(existingBook).Error; err != nil {
		return nil, err
	}

	return existingBook, nil

}

func (r *BookPostgresRepository) DeleteBook(id uint) error {

	return r.db.Delete(&domain.Book{}, id).Error

}
