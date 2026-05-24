package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {

	auth := router.Group("/auth")

	auth.POST("/register", RegisterHandler)
	auth.POST("/login", LoginHandler)
	auth.POST("/logout", LogoutHandler)
	auth.POST("/verifytoken", AuthMiddleware(), VerifyAccessToken)
}
