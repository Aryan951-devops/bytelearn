package auth

import "github.com/gin-gonic/gin"

// RegisterRoutes defines the HTTP routes for the auth package.
func RegisterRoutes(router *gin.RouterGroup) {

	auth := router.Group("/auth")

	auth.POST("/register", RegisterHandler)
	auth.POST("/login", LoginHandler)
	auth.POST("/logout", LogoutHandler)
	auth.POST("/verifytoken", Middleware(), VerifyAccessToken)
}
