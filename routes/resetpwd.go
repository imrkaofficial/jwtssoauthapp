package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/imrkaofficial/jwtssoauthapp/handlers"
)

func RegisterResetPasswordRoutes(router *gin.RouterGroup) {
	router.POST("/reset-password", handlers.ResetPassword)
}
