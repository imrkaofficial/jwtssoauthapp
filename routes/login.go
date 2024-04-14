package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func SetupLoginRoutes(router *gin.Engine) {
	login := router.Group("/login")
	{
		login.GET("/", handlers.ShowLoginPage)
		login.POST("/", handlers.Login)
	}
}