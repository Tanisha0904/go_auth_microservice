package main

import (
	"go-jwt-auth/internal/auth"
	"go-jwt-auth/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//public routes
	r.POST("/login", auth.LoginHandler)

	//protected routes(require jwt)
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/home", func(ctx *gin.Context) {
			username, _ := ctx.Get("username")
			ctx.JSON(200, gin.H{"message": "Welcome to Home", "user": username})
		})
	}
	r.Run(":8080")

}
