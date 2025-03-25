package handlers

import (
	"auth-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{} // Simule une base de données

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	users = append(users, user)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for _, user := range users {
		if user.Email == credentials.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) == nil {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}
