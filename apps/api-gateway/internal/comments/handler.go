package comments

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllCommentOfVideoHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	comments, err := GetAllCommentOfVideoService(videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"comments fetched successfully",
		gin.H{
			"comments": comments,
		},
	))
}

func CreateCommentHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	var req CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
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

	createdComment, err := CreateCommentService(req, user.ID, videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"comment created successfully",
		gin.H{
			"comment": createdComment,
		},
	))
}

func DeleteCommentHandler(c *gin.Context) {

	commentID, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid comment ID format",
			nil,
		))
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

	err = DeleteCommentService(
		commentID,
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"comment deleted successfully",
		nil,
	))
}

func UpdateCommentHandler(c *gin.Context) {

	commentID, err := uuid.Parse(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid comment ID format",
			nil,
		))
		return
	}

	var req CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
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

	updatedComment, err := UpdateCommentService(req, user.ID, commentID)

	if err != nil {
		c.JSON(http.StatusForbidden, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"comment updated successfully",
		gin.H{
			"comment": updatedComment,
		},
	))
}
