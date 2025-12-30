package auth

import (
	"fmt"
	"go-jwt-auth/internal/models"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_SECRET")
var jwtKeyExpirationHour = os.Getenv("JWT_KEY_EXPIRATION_HOUR")

func GenerateToken(username string) (string, error) {
	// 1. Convert string to int
	hours, err := strconv.Atoi(jwtKeyExpirationHour)
	if err != nil {
		// Handle the error if the env variable isn't a valid number
		log.Printf("Invalid expiration hour: %v. Defaulting to 24.", err)
		hours = 24
	}
	//Create a new claim, sign it with a SecretKey using the HMAC-SHA256 algorithm.
	expirationTime := time.Now().Add(time.Duration(hours) * time.Hour)

	claims := &models.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 2. Ensure jwtKey is []byte
	// If jwtKey is a string at the top of your file, cast it here:
	signedToken, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		fmt.Printf("Error signing token: %v\n", err)
		return "", err
	}

	return signedToken, nil
}
