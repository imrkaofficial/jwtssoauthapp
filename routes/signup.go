package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func RegisterSignupRoutes(router *gin.RouterGroup) {
	router.GET("/signup", handlers.ShowSignupPage)
	router.POST("/signup", handlers.Signup)
}
