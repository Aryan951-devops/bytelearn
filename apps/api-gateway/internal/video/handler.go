package video

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetAllVideosHandler(c *gin.Context) {

	videos, err := GetAllVideos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"all videos fetched successfully",
		gin.H{
			"videos": videos,
		},
	))

}
