package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pmas98/go-architecture/modular-architecture/internal/controllers"
	middleware "github.com/pmas98/go-architecture/modular-architecture/internal/middlewares"
)

func RegisterBookRoutes(router *gin.Engine, bookController *controllers.BookController, secretKey string) {
	bookRoutes := router.Group("/books")
	{
		bookRoutes.Use(middleware.AuthMiddleware(secretKey))

		bookRoutes.GET("/", bookController.ListAvailableBooks)
		bookRoutes.GET("/:id", bookController.GetBook)
		bookRoutes.POST("/", bookController.AddBook)
		bookRoutes.POST("/rent/:id", bookController.RentBook)
		bookRoutes.POST("/return/:id", bookController.ReturnBook)
		bookRoutes.PUT("/:id", bookController.UpdateBook)
		bookRoutes.DELETE("/:id", bookController.DeleteBook)
	}
}
