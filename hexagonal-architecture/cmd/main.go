package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/pmas98/go-architecture/hexagonal-architecture/internal/adapters"
	cases "github.com/pmas98/go-architecture/hexagonal-architecture/internal/application/usecases"
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/entity"
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/repositories"
	"github.com/pmas98/go-architecture/hexagonal-architecture/internal/domain/services"
	config "github.com/pmas98/go-architecture/hexagonal-architecture/internal/infrastructure/config"
	infrastructure "github.com/pmas98/go-architecture/hexagonal-architecture/internal/infrastructure/db"
)

func main() {
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, dbConnectionString := config.LoadConfig()

	dbConn := infrastructure.NewDatabaseConnection(dbConnectionString)
	dbConn.AutoMigrate(&entity.Book{}, &entity.User{})

	bookRepo := repositories.NewBookRepository(dbConn)
	bookService := services.NewBookService(bookRepo)
	rentBookUseCase := cases.NewBookUseCases(bookService)
	bookHandler := handlers.NewBookController(rentBookUseCase)
	r.POST("/books/:id/rent", bookHandler.RentBook)

	r.Run() // listen and serve on 0.0.0.0:8080
}
