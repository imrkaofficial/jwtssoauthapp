package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowSignupPage renders the signup page
func ShowSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// Signup handles user signup
func Signup(c *gin.Context) {
	// Handle signup logic
}