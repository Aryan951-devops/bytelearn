package like

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ToggleCommentLikeHandler(c *gin.Context) {

	commentID, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			utils.NewResponse("invalid comment id", nil))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	liked, err := ToggleCommentLikeService(
		user.ID,
		commentID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"comment like toggled",
			gin.H{
				"liked": liked,
			},
		))
}

func ToggleVideoLikeHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			utils.NewResponse("invalid video id", nil))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	liked, err := ToggleVideoLikeService(
		user.ID,
		videoID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"video like toggled",
			gin.H{
				"liked": liked,
			},
		))
}

func CheckVideoLikeHandler(c *gin.Context) {

	videoID, _ := uuid.Parse(c.Param("videoId"))

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	liked, err := CheckVideoLikeService(
		user.ID,
		videoID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"video like status fetched",
			gin.H{
				"liked": liked,
			},
		))
}

func CheckCommentLikeHandler(c *gin.Context) {

	commentID, _ := uuid.Parse(c.Param("commentId"))

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	liked, err := CheckCommentLikeService(
		user.ID,
		commentID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"comment like status fetched",
			gin.H{
				"liked": liked,
			},
		))
}

func GetTotalVideoLikesHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			utils.NewResponse("invalid video id", nil))
		return
	}

	total, err := GetTotalVideoLikesService(videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"video likes count fetched",
			gin.H{
				"total_likes": total,
			},
		))
}

func GetTotalCommentLikesHandler(c *gin.Context) {

	commentID, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest,
			utils.NewResponse("invalid comment id", nil))
		return
	}

	total, err := GetTotalCommentLikesService(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.NewResponse(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK,
		utils.NewResponse(
			"comment likes count fetched",
			gin.H{
				"total_likes": total,
			},
		))
}
