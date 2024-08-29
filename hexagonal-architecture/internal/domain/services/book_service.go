package services

import (
	"errors"

	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/entity"
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/repositories"
)

type BookService struct {
	repo *repositories.BookRepository
}

func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) ListBooks() ([]*entity.Book, error) {
	return s.repo.GetAllBooks()
}

func (s *BookService) AddBook(book *entity.Book) (uint, error) {
	if book.Title == "" || book.Author == "" {
		return 0, errors.New("title and author cannot be empty")
	}
	if err := s.validateBook(book); err != nil {
		return 0, err
	}
	return s.repo.CreateBook(book)
}

func (s *BookService) GetBook(id uint) (*entity.Book, error) {
	if id == 0 {
		return nil, errors.New("invalid book ID")
	}
	return s.repo.GetBookByID(id)
}

func (s *BookService) UpdateBook(book *entity.Book) error {
	if book.ID == 0 {
		return errors.New("book ID cannot be zero")
	}
	if err := s.validateBook(book); err != nil {
		return err
	}
	return s.repo.UpdateBook(book)
}

func (s *BookService) DeleteBook(id uint) error {
	if id == 0 {
		return errors.New("invalid book ID")
	}
	return s.repo.DeleteBook(id)
}

func (s *BookService) validateBook(book *entity.Book) error {
	existingBook, _ := s.repo.FindByTitle(book.Title)
	if existingBook != nil && existingBook.ID != book.ID {
		return errors.New("a book with this title already exists")
	}
	return nil
}
