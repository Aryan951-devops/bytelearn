package user

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	user := router.Group("/user")
	user.Use(auth.AuthMiddleware())

	user.POST("/change-password", ChangePasswordHandler)
	user.GET("/current-user", GetUserHandler)
	user.POST("/update-account", UpdateAccountHandler)
}
