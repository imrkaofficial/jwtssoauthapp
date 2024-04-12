package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/imrkaofficial/jwtssoauthapp/models"
)

// ShowLoginPage renders the login page
func ShowLoginPage(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", nil)
}

// Login handles user login
func Login(c *gin.Context) {
}

// ShowSignupPage renders the signup page
func ShowSignupPage(c *gin.Context) {
    c.HTML(http.StatusOK, "signup.html", nil)
}

// Signup handles user signup
func Signup(c *gin.Context) {
}

// ShowForgotPasswordPage renders the forgot password page
func ShowForgotPasswordPage(c *gin.Context) {
    c.HTML(http.StatusOK, "forgotpwd.html", nil)
}

// ForgotPassword handles the forgot password logic
func ForgotPassword(c *gin.Context) {
}

// ResetPassword handles the password reset logic
func ResetPassword(c *gin.Context) {
}
