package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func SetupResetPasswordRoutes(router *gin.Engine) {
	ResetPassword := router.Group("/reset-password")
	{
		ResetPassword.POST("/", handlers.ResetPassword)
	}
}
