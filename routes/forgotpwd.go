package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func SetupForgotPasswordRoutes(router *gin.Engine) {
	ForgotPassword := router.Group("/forgotpwd")
	{
		ForgotPassword.GET("/", handlers.ShowForgotPasswordPage)
		ForgotPassword.POST("/", handlers.ForgotPassword)
	}
}
