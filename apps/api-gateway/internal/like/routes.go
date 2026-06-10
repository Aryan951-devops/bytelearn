package like

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the like package.
func RegisterRoutes(router *gin.RouterGroup) {

	like := router.Group("/like")

	like.POST("/toggle/c/:commentId", auth.Middleware(), ToggleCommentLikeHandler)
	like.POST("/toggle/v/:videoId", auth.Middleware(), ToggleVideoLikeHandler)
	like.GET("/check/c/:commentId", auth.Middleware(), CheckCommentLikeHandler)
	like.GET("/check/v/:videoId", auth.Middleware(), CheckVideoLikeHandler)
	like.GET("/total/v/:videoId", GetTotalVideoLikesHandler)
	like.GET("/total/c/:commentId", GetTotalCommentLikesHandler)

}
