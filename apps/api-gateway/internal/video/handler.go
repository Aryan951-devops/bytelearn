package video

import (
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"message": "all videos fetched succesfully",
		"videos":  videos,
	})

}
