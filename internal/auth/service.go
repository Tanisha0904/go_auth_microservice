package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = os.Getenv("JWT_SECRET")

func GenerateToken(username string) (string, error) {
	//Create a new claim, sign it with a SecretKey using the HMAC-SHA256 algorithm.
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
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
