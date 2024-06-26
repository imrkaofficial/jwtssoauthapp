package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/config"
	"github.com/imrkaofficial/jwtssoauthapp/routes"
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
	routes.SetupLoginRoutes(router)
	routes.SetupSignupRoutes(router)
	routes.SetupForgotPasswordRoutes(router)
	routes.SetupResetPasswordRoutes(router)

	fmt.Println("Server running on localhost:9000")
	router.Run(":9000")
}