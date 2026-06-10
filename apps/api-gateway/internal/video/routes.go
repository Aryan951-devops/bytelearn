package video

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the video package.
func RegisterRoutes(router *gin.RouterGroup) {

	video := router.Group("/video")

	video.GET("/", auth.Middleware(), GetAllVideosOfUserHandler)
	video.GET("/all", GetAllVideosHandler)
	video.GET("/upload/signature", auth.Middleware(), GenerateUploadSignatureHandler)
	video.POST("/upload", auth.Middleware(), UploadVideoHandler)
	video.GET("/:videoID", GetVideoHandler)
	video.PATCH("/:videoID", auth.Middleware(), UpdateVideoHandler)
	video.DELETE("/:videoID", auth.Middleware(), DeleteVideoHandler)

}
