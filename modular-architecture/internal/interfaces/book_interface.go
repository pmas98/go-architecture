package interfaces

import "github.com/pmas98/go-architecture/modular-architecture/internal/domain"

type BookRepository interface {
	GetBookByID(id uint) (*domain.Book, error)
	FindByTitle(title string) (*domain.Book, error)
	CreateBook(book *domain.Book) (uint, error)
	UpdateBook(id uint, book *domain.Book) error
	DeleteBook(id uint) error
}
