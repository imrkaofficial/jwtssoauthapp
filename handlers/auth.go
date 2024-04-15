package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/config"
	"github.com/imrkaofficial/jwtssoauthapp/models"
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

// // Signup handles user signup
// func Signup(c *gin.Context) {
// }

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

// Signup handles user signup
func Signup(c *gin.Context) {
	// Bind the form data
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Validate the data (for example, you can use a validation library)
	if err := validateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username or email already exists
	if userExists := checkUserExists(user.Username, user.Email); userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exists"})
		return
	}

	// Hash the password before saving
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword

	// Save the user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	// Send success response
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// validateUser validates the user data (you can customize this function)
func validateUser(user models.User) error {
	// Example validation, you can implement your own validation rules
	if len(user.Username) < 4 {
		return fmt.Errorf("username must be at least 4 characters")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}
	return nil
}

// checkUserExists checks if the username or email already exists
func checkUserExists(username, email string) bool {
	var count int64
	config.DB.Model(&models.User{}).Where("username = ? OR email = ?", username, email).Count(&count)
	return count > 0
}

// hashPassword hashes the password before saving to the database
func hashPassword(password string) (string, error) {
	// Example hashing function, you should use a secure hashing algorithm
	// For example, bcrypt: https://pkg.go.dev/golang.org/x/crypto/bcrypt
	return password, nil
}