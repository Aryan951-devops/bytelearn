package like

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	like := router.Group("/like")

	like.POST("/toggle/c/:commentId", auth.AuthMiddleware(), ToggleCommentLikeHandler)
	like.POST("/toggle/v/:videoId", auth.AuthMiddleware(), ToggleVideoLikeHandler)
	like.GET("/check/c/:commentId", auth.AuthMiddleware(), CheckCommentLikeHandler)
	like.GET("/check/v/:videoId", auth.AuthMiddleware(), CheckVideoLikeHandler)
	like.GET("/total/v/:videoId", GetTotalVideoLikesHandler)
	like.GET("/total/c/:commentId", GetTotalCommentLikesHandler)

}
