package domain

type Book struct {
	ID              uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Title           string `json:"title" gorm:"type:varchar(255);not null"`
	Author          string `json:"author" gorm:"type:varchar(255);not null"`
	PublicationYear int    `json:"publication_year" gorm:"not null"`
}

type BookRepository interface {
	CreateBook(book *Book) (*Book, error)
	ReadBooks() ([]*Book, error)
	ReadBook(id uint) (*Book, error)
	UpdateBook(id uint, book *Book) (*Book, error)
	DeleteBook(id uint) error
}
