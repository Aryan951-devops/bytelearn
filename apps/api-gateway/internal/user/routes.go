package user

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the user package.
func RegisterRoutes(router *gin.RouterGroup) {

	user := router.Group("/user")
	user.Use(auth.Middleware())

	user.POST("/change-password", ChangePasswordHandler)
	user.GET("/current-user", GetUserHandler)
	user.PATCH("/update-account", UpdateAccountHandler)
}
