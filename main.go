package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Route de santé
	router.GET("/567088a9-6689-4e67-b5e5-ed40ad0a830c/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API is running!"})
	})

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
