package pkg

const (
	ErrNotFound       = "book not found"
	ErrInvalidID      = "invalid book ID"
	ErrInvalidRequest = "invalid request payload"
)

var Response = map[string]map[string]string{
	ErrNotFound:       {"error": "Book not found"},
	ErrInvalidID:      {"error": "Invalid book ID"},
	ErrInvalidRequest: {"error": "Invalid request payload"},
}
