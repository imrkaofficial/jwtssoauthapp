package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowLoginPage renders the login page
func ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// Login handles user login
func Login(c *gin.Context) {
	// Handle login logic
}

