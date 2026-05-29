package video

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllVideosHandler(c *gin.Context) {

	videos, err := GetAllVideos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"all videos fetched successfully",
		gin.H{
			"videos": videos,
		},
	))
}

func GetAllVideosOfUserHandler(c *gin.Context) {

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	videos, err := GetAllVideosOfUserService(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"user's all videos are fetched successfully",
		gin.H{
			"videos": videos,
		},
	))
}

func GetVideoHandler(c *gin.Context) {

	videoId, err := uuid.Parse(c.Param("videoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	video, err := GetVideo(videoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"video fetched successfully",
		gin.H{
			"video": video,
		},
	))
}

func DeleteVideoHandler(c *gin.Context) {

	videoId, err := uuid.Parse(c.Param("videoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
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

	video, err := DeleteVideo(videoId, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"video deleted successfully",
		gin.H{
			"video": video,
		},
	))
}

func GenerateUploadSignatureHandler(c *gin.Context) {

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)
	if user.Role != "educator" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are not allowed to get signature",
			nil,
		))
		return
	}

	signatureData, err := GenerateUploadSignature()

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"video signature generated successfully",
		signatureData,
	))
}

func UploadVideoHandler(c *gin.Context) {

	var req UploadVideoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request body",
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
	if user.Role != "educator" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are not allowed to upload videos",
			nil,
		))
		return
	}

	video, err := UploadVideo(req, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			"failed to create video",
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"video uploaded successfully",
		gin.H{
			"video": video,
		},
	))
}

func UpdateVideoHandler(c *gin.Context) {

	videoId, err := uuid.Parse(c.Param("videoID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	var req UpdateVideoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	file, err := c.FormFile("thumbnail")

	var updated_video *models.Video

	if err == nil {
		if err := utils.IsImageAllowed(file); err != nil {
			c.JSON(http.StatusBadRequest, utils.NewResponse(
				err.Error(),
				nil,
			))
			return
		}

		os.MkdirAll("./temp", os.ModePerm)

		// unique file name
		tempPath := fmt.Sprintf(
			"./temp/%s%s",
			uuid.New().String(),
			filepath.Ext(file.Filename),
		)

		if err := c.SaveUploadedFile(file, tempPath); err != nil {
			c.JSON(http.StatusInternalServerError, utils.NewResponse(
				"failed to save uploaded file",
				nil,
			))
			return
		}
		defer os.Remove(tempPath)

		updated_video, err = UpdateVideo(req, user.ID, videoId, &tempPath)
	} else {
		updated_video, err = UpdateVideo(req, user.ID, videoId, nil)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"video updated successfully",
		gin.H{
			"user": updated_video,
		},
	))
}
