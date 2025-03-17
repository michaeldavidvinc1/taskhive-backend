package routes

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(middleware.LoggerMiddleware())

	// Public routes
	//public := r.Group("/api")
	//public.POST("/register", controller.Register)
	//public.POST("/login", controller.Login)

	// Protected routes
	//protected := r.Group("/api")
	//protected.Use(middleware.JWTAuthMiddleware())
	//{
	//	// Add protected routes here
	//	protected.GET("/profile", controller.Profile)
	//}
	return r
}
