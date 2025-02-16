package routes

import (
	"seaventures/src/controller"
	"seaventures/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controller.RegisterUser)
		userRoutes.POST("/login", controller.LoginUser)

	}


	authRoutes := r.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.POST("/protected", controller.ProtectedEndpoint)
	}
}

