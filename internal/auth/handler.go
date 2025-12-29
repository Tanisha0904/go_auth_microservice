package auth

import (
	"fmt"
	"go-jwt-auth/internal/database"
	"go-jwt-auth/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Handler: Receives JSON, validates the "password" (keep it hardcoded for now), calls the GenerateToken service, and returns the JWT to the user.
func LoginHandler(c *gin.Context) {
	var req models.LoginRequest
	var user models.User

	// 1. Bind JSON input to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//Query DB for the user
	result := database.DB.Where("username = ?", req.Username).First(&user)

	if result.Error != nil {
		c.JSON(401, gin.H{"error": "User not found, please sign in first"})
		return
	}
	// 2. Authentication (Mocking a database check)
	// In a real app, use bcrypt.CompareHashAndPassword here
	if req.Password != user.Password {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}
	token, err := GenerateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

func SignUp(c *gin.Context) {
	fmt.Println("\n\n into signup function")
	var req models.SignUpRequest

	// 1. Bind JSON input to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// 2. Check if user already exists
	var existingUser models.User
	// .First() executes the query. If it finds a record, err will be nil.
	err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error

	if err == nil {
		// If no error, it means a user was found
		c.JSON(http.StatusConflict, gin.H{"error": "User already present, please login"})
		return
	}

	// 3. Validate passwords
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password didnt match"})
		c.Abort()
		return
	}

	// 4. Create the user (Initialize here so req.Username has data)
	user := models.User{Username: req.Username, Password: req.Password}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}
