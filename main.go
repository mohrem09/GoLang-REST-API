package main

import (
	"auth-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.AuthRoutes(router)

	router.Run(":8080")
}
