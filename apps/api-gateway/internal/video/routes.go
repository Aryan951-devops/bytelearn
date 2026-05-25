package video

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	video := router.Group("/video")

	video.GET("/all", GetAllVideosHandler)
}
