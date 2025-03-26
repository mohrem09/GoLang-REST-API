package routes

import (
	"auth-api/handlers"
	"auth-api/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	api := router.Group("/mohamed-rizwane")
	{
		router.GET("/health", func(c *gin.Context) {
			c.Status(200)
		})
		router.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Bienvenue sur l'API d'authentification!",
			})
		})

		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/me", middlewares.AuthMiddleware(), handlers.Profile)

	}
}
