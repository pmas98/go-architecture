package repositories

import (
	"log"

	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/entity"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) GetAllBooks() ([]*entity.Book, error) {
	var books []*entity.Book

	// Log the start of the function
	log.Println("Getting all books")

	// Use Preload if there are any relationships to load
	result := r.db.Preload("RentedBy").Find(&books)

	// Log the result
	log.Printf("Found %d books", len(books))

	if result.Error != nil {
		log.Printf("Error getting books: %v", result.Error)
		return nil, result.Error
	}

	// If no books were found, log this information
	if len(books) == 0 {
		log.Println("No books found in the database")
	}

	return books, nil
}

func (r *BookRepository) GetBookByID(id uint) (*entity.Book, error) {
	var book entity.Book
	result := r.db.Preload("RentedBy").First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) FindByTitle(title string) (*entity.Book, error) {
	var book entity.Book
	result := r.db.Where("title = ?", title).First(&book)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // No book found
		}
		return nil, result.Error // Other errors
	}
	return &book, nil
}

// Add a new book to the database
func (r *BookRepository) CreateBook(book *entity.Book) (uint, error) {
	result := r.db.Create(book)
	if result.Error != nil {
		return 0, result.Error
	}
	return book.ID, nil
}

// Update an existing book in the database
func (r *BookRepository) UpdateBook(book *entity.Book) error {
	result := r.db.Save(book)
	return result.Error
}

// Delete a book from the database
func (r *BookRepository) DeleteBook(id uint) error {
	result := r.db.Delete(&entity.Book{}, id)
	return result.Error
}
