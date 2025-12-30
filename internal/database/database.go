// This function initializes the connection and performs Auto-Migration
package database

import (
	"fmt"
	"go-jwt-auth/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres_db_user"
// 	password = "postgres_123"
// 	dbname   = "go_auth_db"
// )

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_auth_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//Auto-create the users table
	db.AutoMigrate((&models.User{}))
	DB = db
	fmt.Println("Databse connection successful")

}
