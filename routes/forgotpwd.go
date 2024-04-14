package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func RegisterForgotPasswordRoutes(router *gin.RouterGroup) {
	router.GET("/forgotpwd", handlers.ShowForgotPasswordPage)
	router.POST("/forgotpwd", handlers.ForgotPassword)
}
