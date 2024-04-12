package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
	"github.com/imrkaofficial/jwtssoauthapp/config"
)

func main() {
	// Connect to MySQL database
	dsn := "admin:jwtsso!!1234@tcp(localhost:3306)/jwtssodb?charset=utf8mb4&parseTime=True&loc=Local"
	config.Connect(dsn) // Initialize and connect to the database

	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Routes
	router.GET("/login", handlers.ShowLoginPage)
	router.POST("/login", handlers.Login)
	router.GET("/signup", handlers.ShowSignupPage)
	router.POST("/signup", handlers.Signup)
	router.GET("/forgotpwd", handlers.ShowForgotPasswordPage)
	router.POST("/forgotpwd", handlers.ForgotPassword)
	router.POST("/reset-password", handlers.ResetPassword)

	fmt.Println("Server running on localhost:9000")
	router.Run(":9000")
}