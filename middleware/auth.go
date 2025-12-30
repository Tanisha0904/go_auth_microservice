package middleware

import (
	"fmt"
	"go-jwt-auth/internal/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// This is a "Gatekeeper" function. It intercepts incoming requests, looks for the Authorization header, parses the JWT, and decides if the request can proceed.
//It follows the Extract -> Parse -> Validate pattern.

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//1. extract header
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			ctx.Abort()
			return
		}

		//2. Parse Bearer token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		//3. Validate token
		if err != nil {
			fmt.Println("JWT Parse Error:", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}
		if !token.Valid {
			fmt.Print("\n\ntoken in auth: ", token)
		}
		//4. set context(pawss user data to subsequent handlers)
		ctx.Set("username", claims.Username)
		ctx.Next()
	}
}
