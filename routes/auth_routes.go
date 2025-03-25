package routes

import (
	"auth-api/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	api := router.Group("/567088a9-6689-4e67-b5e5-ed40ad0a830c")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}
}
