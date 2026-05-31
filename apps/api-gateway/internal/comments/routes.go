package comments

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	comment := router.Group("/comment")

	comment.GET("/v/:videoId", GetAllCommentOfVideoHandler)
	comment.POST("/v/:videoId", auth.AuthMiddleware(), CreateCommentHandler)
	comment.DELETE("/c/:commentId", auth.AuthMiddleware(), DeleteCommentHandler)
	comment.PATCH("/c/:commentId", auth.AuthMiddleware(), UpdateCommentHandler)

}
