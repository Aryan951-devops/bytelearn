package video

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	video := router.Group("/video")

	video.GET("/upload/signature", auth.AuthMiddleware(), GenerateUploadSignatureHandler)
	video.POST("/upload", auth.AuthMiddleware(), UploadVideoHandler)
	video.GET("/:videoID", GetVideoHandler)
	video.PATCH("/:videoID", auth.AuthMiddleware(), UpdateVideoHandler)
	video.DELETE("/:videoID", auth.AuthMiddleware(), DeleteVideoHandler)
	video.GET("/all", GetAllVideosHandler)

}
