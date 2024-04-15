package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowForgotPasswordPage renders the forgot password page
func ShowForgotPasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "forgotpwd.html", nil)
}

// ForgotPassword handles the forgot password logic
func ForgotPassword(c *gin.Context) {
	// Handle forgot password logic
}