package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"go-jwt-auth/internal/database"
	"go-jwt-auth/internal/models"
	"net/http"
	"time"

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
	user := models.User{Username: req.Username, Password: req.Password, Email: req.Email}
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	var user models.User
	err := database.DB.Where("email=?", req.Email).First(&user).Error
	if err != nil {
		// If error, it means a user was not found
		c.JSON(http.StatusConflict, gin.H{"error": "User not found, please signup first"})
		return
	}
	//1.Generate a sceure random token
	b := make([]byte, 16)
	//Use crypto/rand for unguessable tokens (don't use JWT for this).
	rand.Read(b)
	token := hex.EncodeToString(b)

	// 2. Save token and expiry (15 mins from now) to DB
	user.PasswordResetToken = token
	user.PasswordResetAt = time.Now().Add(5 * time.Minute)
	database.DB.Save(&user)

	//3. construct the reset url
	// this url will point to frontend url
	reset_url := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)

	c.JSON(http.StatusOK, gin.H{
		"message":   "Password reset link generated",
		"reset_url": reset_url,
	})

}

func ResetPassword(c *gin.Context) {
	// Get token from URL query: /reset-password?token=abcdef...
	token := c.Query(("token"))
	fmt.Println("\n token: ", token)

	var req models.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}
	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password and confirm password didnt match"})
		c.Abort()
		return
	}
	//Find user by token and ensure the token hasn't expired
	var user models.User
	err := database.DB.Where("password_reset_token = ? AND password_reset_at > ?", token, time.Now()).First(&user).Error
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Invalid or expired token, please try reseting password again"})
		return
	}

	//update password and clear the token fields
	user.Password = req.Password // IMPORTANT: Hash this with bcrypt before saving!
	user.PasswordResetToken = ""
	user.PasswordResetAt = time.Time{}
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Password has been reset successfully"})
}
