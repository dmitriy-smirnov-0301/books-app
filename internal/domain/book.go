package domain

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publication_year"`
}

type BookRepository interface {
	CreateBook(book *Book) (*Book, error)
	ReadBooks() ([]*Book, error)
	ReadBook(id int) (*Book, error)
	UpdateBook(id int, book *Book) (*Book, error)
	DeleteBook(id int) error
}
