package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pmas98/go-architecture/modular-architecture/config"
	"github.com/pmas98/go-architecture/modular-architecture/internal/controllers"
	"github.com/pmas98/go-architecture/modular-architecture/internal/domain"
	"github.com/pmas98/go-architecture/modular-architecture/internal/infrastructure"
	"github.com/pmas98/go-architecture/modular-architecture/internal/repositories"
	"github.com/pmas98/go-architecture/modular-architecture/internal/routes"
	"github.com/pmas98/go-architecture/modular-architecture/internal/services"
	cases "github.com/pmas98/go-architecture/modular-architecture/internal/usecases"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	secretKey, dbConnectionString := config.LoadConfig()

	db := infrastructure.NewDatabaseConnection(dbConnectionString)

	db.AutoMigrate(&domain.Book{}, &domain.User{})

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userUseCase := cases.NewUserUseCase(userService, secretKey)

	bookRepo := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepo)
	bookCase := cases.NewBookUseCases(bookService)
	bookController := controllers.NewBookController(bookCase)
	userController := controllers.NewUserController(userUseCase)

	router := gin.Default()

	routes.RegisterBookRoutes(router, bookController, secretKey)
	routes.RegisterUserRoutes(router, userController)

	router.Run(":8080")
}
