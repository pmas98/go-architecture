package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	cases "github.com/pmas98/go-architecture/modular-architecture/internal/usecases"
)

type BookController struct {
	useCase *cases.BookUseCases
}

func NewBookController(usecase *cases.BookUseCases) *BookController {
	return &BookController{useCase: usecase}
}

func (c *BookController) AddBook(ctx *gin.Context) {
	var book domain.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := c.useCase.AddBook(book.Title, book.Author, book.PublishedAt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	book.ID = uint(id)
	ctx.JSON(http.StatusCreated, book)
}

func (c *BookController) GetBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	book, err := c.useCase.GetBookDetails(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	bookDetails := map[string]interface{}{
		"ID":          book.ID,
		"Title":       book.Title,
		"Author":      book.Author,
		"PublishedAt": book.PublishedAt,
		"IsAvailable": book.IsAvailable,
		"RentedByID":  book.RentedByID,
	}

	if book.RentedBy != nil {
		bookDetails["RentedByUser"] = book.RentedBy.Username
	} else {
		bookDetails["RentedByUser"] = nil
	}

	ctx.JSON(http.StatusOK, bookDetails)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	var book domain.Book
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := c.useCase.UpdateBook(uint(id), book.Title, book.Author, book.PublishedAt); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := c.useCase.DeleteBook(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *BookController) ReturnBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	err_return := c.useCase.ReturnBook(uint(id), userIDUint)
	if err_return != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Only the user who rented can return"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"returned": "true"})
}

func (c *BookController) RentBook(ctx *gin.Context) {
	bookID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	err = c.useCase.RentBook(uint(bookID), userIDUint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book rented successfully"})
}

func (c *BookController) ListAvailableBooks(ctx *gin.Context) {

	books, err := c.useCase.ListAvailableBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"books": books})
}
