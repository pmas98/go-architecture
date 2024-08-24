package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pmas98/go-architecture/modular-architecture/internal/controllers"
)

func RegisterUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
}
