package handlers

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/config"
	"github.com/imrkaofficial/jwtssoauthapp/models"
	"golang.org/x/crypto/bcrypt"
)

// ShowSignupPage renders the signup page
func ShowSignupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

// GetUserByID retrieves user data by ID
func GetUserByID(ctx *gin.Context) {
    // Get user ID from path parameter
    userID := ctx.Param("id")

    // Query the database for the user
    var users models.User
    if err := config.DB.First(&users, userID).Error; err != nil {
        // If user data not found, return 404 Not Found error
        ctx.JSON(404, gin.H{"error": "User not found"})
        return
    }

    // If user data found, return it in the response
    ctx.JSON(200, users)
}

// GetAllUsers retrieves all user data
func GetAllUsers(ctx *gin.Context) {
    var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(200, users)
}


func CreateUser(ctx *gin.Context) {
	var users models.User
	if err := ctx.BindJSON(&users); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request data"})
		return
	}
	if err := config.DB.Create(&users).Error; err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to create account"})
		return
	}
	ctx.JSON(201, gin.H{"message": "User has been created Successfully"})
}


// Signup handles user signup
func Signup(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"Error":   "Invalid form data",
			"Success": "",
		})
		return
	}

	// Validate form fields
	if err := validateSignupForm(form); err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"Error":   err.Error(),
			"Success": "",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{
			"Error":   "Error hashing password",
			"Success": "",
		})
		return
	}

	// Create user
	newUser := models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		Username:  form.Username,
		Email:     form.Email,
		Password:  string(hashedPassword),
	}

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{
			"Error":   "Error creating user",
			"Success": "",
		})
		return
	}

	c.HTML(http.StatusOK, "signup.html", gin.H{
		"Error":   "",
		"Success": "User created successfully",
	})
}

func validateSignupForm(form models.User) error {
	// Validate First Name
	if !isValidName(form.FirstName) {
		return fmt.Errorf("please enter valid first name")
	}

	// Validate Last Name
	if !isValidName(form.LastName) {
		return fmt.Errorf("please enter valid last name")
	}

	// Validate Email
	if !isValidEmail(form.Email) {
		return fmt.Errorf("please enter valid email")
	}

	// Validate Password
	if !isValidPassword(form.Password) {
		return fmt.Errorf("password must contain at least one number, one uppercase letter, one lowercase letter, and be at least 8 characters long")
	}

	// Validate Confirm Password
	// if form.Password != form.ConfPwd {
	// 	return fmt.Errorf("Confirm Password does not match")
	// }

	// Validate Username
	if !isValidUsername(form.Username) {
		return fmt.Errorf("username must be alphanumeric and at least 6 characters long")
	}

	return nil
}

func isValidName(name string) bool {
	// Name should only contain letters, spaces, and dots
	// Minimum 2 characters
	validName := regexp.MustCompile(`^[a-zA-Z .]{2,}$`)
	return validName.MatchString(name)
}

func isValidEmail(email string) bool {
	// Email regex pattern
	emailPattern := `^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$`
	validEmail := regexp.MustCompile(emailPattern)
	return validEmail.MatchString(email)
}

func isValidPassword(password string) bool {
	// Password should contain at least one number, one uppercase letter, one lowercase letter
	// and be at least 8 characters long
	containsNumber := false
	containsUpper := false
	containsLower := false

	for _, char := range password {
		switch {
		case char >= '0' && char <= '9':
			containsNumber = true
		case char >= 'A' && char <= 'Z':
			containsUpper = true
		case char >= 'a' && char <= 'z':
			containsLower = true
		}
	}
	return containsNumber && containsUpper && containsLower && len(password) >= 8
}

func isValidUsername(username string) bool {
	// Username should be alphanumeric and at least 6 characters
	validUsername := regexp.MustCompile(`^[a-zA-Z0-9]{6,}$`)
	return validUsername.MatchString(username)
}
