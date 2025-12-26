package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Handler: Receives JSON, validates the "password" (keep it hardcoded for now), calls the GenerateToken service, and returns the JWT to the user.
func LoginHandler(c *gin.Context) {
	var req LoginRequest

	// 1. Bind JSON input to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// 2. Authentication (Mocking a database check)
	// In a real app, use bcrypt.CompareHashAndPassword here
	if req.Username == "admin" && req.Password == "password_12345" {
		token, err := GenerateToken(req.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Token"})
			return
		}
		fmt.Print("\n\n token in handler: ", token)
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})

}
