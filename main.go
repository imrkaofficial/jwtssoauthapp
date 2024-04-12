package main

import (
	"fmt"
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
	"github.com/imrkaofficial/jwtssoauthapp/models"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	// Connect to MySQL database
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}


	// Auto Migrate the models
	db.AutoMigrate(&models.User{})

	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Pass the DB instance to handlers
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Routes
	router.GET("/login", handlers.ShowLoginPage)
	router.POST("/login", handlers.Login)
	router.GET("/signup", handlers.ShowSignupPage)
	router.POST("/signup", handlers.Signup)
	router.GET("/forgotpwd", handlers.ShowForgotPasswordPage)
	router.POST("/forgotpwd", handlers.ForgotPassword)
	router.POST("/reset-password", handlers.ResetPassword)

	fmt.Println("Server running on localhost:8080")
	router.Run(":8080")
}