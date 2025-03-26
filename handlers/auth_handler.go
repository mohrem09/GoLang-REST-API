package handlers

import (
	"auth-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Convertir en structure User si nécessaire
	userData, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    userData.ID,
		"name":  userData.Name,
		"email": userData.Email,
	})
}
