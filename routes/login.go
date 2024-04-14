package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func RegisterLoginRoutes(router *gin.RouterGroup) {
	router.GET("/login", handlers.ShowLoginPage)
	router.POST("/login", handlers.Login)
}
