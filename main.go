package main

import (
	"auth-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Ajoute ce code pour gérer la racine ("/") et le favicon.ico
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue dans votre API GoLang!",
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(204) // Ignore proprement les requêtes favicon.ico
	})
	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
	// Routes existantes
	api := r.Group("/567088a9-6689-4e67-b5e5-ed40ad0a830c")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/me", handlers.Profile)
	}

	// Lance le serveur sur le port 8080
	r.Run(":8080")
}
