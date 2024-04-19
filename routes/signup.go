package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func SetupSignupRoutes(router *gin.Engine) {
	signup := router.Group("/signup")
	{
		signup.GET("", handlers.ShowSignupPage)
		signup.POST("", handlers.Signup)
		signup.GET("/Users", handlers.GetAllUsers)
		signup.GET("/Users/:id", handlers.GetUserByID)
		signup.POST("/Users", handlers.CreateUser)
	}
}
