package handlers

import (
	"auth-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var users = []models.User{} // Simule une base de données

var jwtKey = []byte("SECRET_KEY") // Remplace par une clé secrète sécurisée

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
			// ✅ Générer un token JWT
			expirationTime := time.Now().Add(24 * time.Hour) // Token valide 24h
			claims := &jwt.RegisteredClaims{
				Subject:   user.Email,
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}

			// ✅ Retourner le token au client
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful",
				"token":   tokenString,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
}
