package cases

import (
	"errors"
	"time"

	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	"github.com/pmas98/go-architecture/modular-architecture/internal/services"
)

type BookUseCases struct {
	bookService *services.BookService
}

func NewBookUseCases(bookService *services.BookService) *BookUseCases {
	return &BookUseCases{bookService: bookService}
}

func (uc *BookUseCases) RentBook(bookID uint, userID uint) error {
	book, err := uc.bookService.GetBook(bookID)
	if err != nil {
		return err
	}

	err = book.Rent(userID)
	if err != nil {
		return err
	}

	return uc.bookService.UpdateBook(book)
}

func (uc *BookUseCases) ReturnBook(bookID uint, userID uint) error {
	book, _ := uc.bookService.GetBook(bookID)

	if *book.RentedByID != userID {
		return errors.New("Cant")
	}

	book.Return()

	return uc.bookService.UpdateBook(book)
}

func (uc *BookUseCases) AddBook(title, author string, publishedAt time.Time) (uint, error) {
	book := &domain.Book{
		Title:       title,
		Author:      author,
		PublishedAt: publishedAt,
		IsAvailable: true,
	}
	return uc.bookService.AddBook(book)
}

func (uc *BookUseCases) UpdateBook(id uint, title, author string, publishedAt time.Time) (uint, error) {
	book := &domain.Book{
		Title:       title,
		Author:      author,
		PublishedAt: publishedAt,
		IsAvailable: true,
	}
	return uc.bookService.AddBook(book)
}

func (uc *BookUseCases) DeleteBook(id uint) error {
	return uc.bookService.DeleteBook(id)
}
func (uc *BookUseCases) GetBookDetails(bookID uint) (*domain.Book, error) {
	return uc.bookService.GetBook(bookID)
}

type RentedByResponse struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
}

type BookResponse struct {
	ID          int               `json:"ID"`
	Title       string            `json:"Title"`
	Author      string            `json:"Author"`
	PublishedAt string            `json:"PublishedAt"`
	IsAvailable bool              `json:"IsAvailable"`
	RentedBy    *RentedByResponse `json:"RentedBy,omitempty"` // Use omitempty to exclude null fields
}

func convertToBookResponse(book *domain.Book) *BookResponse {
	if book == nil {
		return nil
	}

	var rentedByResponse *RentedByResponse
	if !book.IsAvailable && book.RentedBy != nil {
		rentedByResponse = &RentedByResponse{
			Username: book.RentedBy.Username,
			Email:    book.RentedBy.Email,
		}
	}

	return &BookResponse{
		ID:          int(book.ID),
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt.Format("01-02-2006"),
		IsAvailable: book.IsAvailable,
		RentedBy:    rentedByResponse,
	}
}

func (uc *BookUseCases) ListAvailableBooks() ([]*BookResponse, error) {
	books, err := uc.bookService.ListBooks()
	if err != nil {
		return nil, err
	}

	availableBooks := []*BookResponse{}
	for _, book := range books {
		availableBooks = append(availableBooks, convertToBookResponse(book))
	}

	return availableBooks, nil
}
