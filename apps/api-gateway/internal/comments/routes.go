package comments

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the comments package.
func RegisterRoutes(router *gin.RouterGroup) {

	comment := router.Group("/comment")

	comment.GET("/v/:videoId", GetAllCommentOfVideoHandler)
	comment.POST("/v/:videoId", auth.Middleware(), CreateCommentHandler)
	comment.DELETE("/c/:commentId", auth.Middleware(), DeleteCommentHandler)
	comment.PATCH("/c/:commentId", auth.Middleware(), UpdateCommentHandler)

}
