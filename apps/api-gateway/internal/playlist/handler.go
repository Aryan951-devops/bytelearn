package playlist

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePlaylistHandler(c *gin.Context) {

	var req CreatePlaylistRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
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

	if user.Role == "learner" && req.Type == "course" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are authorized to create course",
			nil,
		))
		return
	}

	playlist, err := CreatePlaylistService(req, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"playlist created successfully",
		gin.H{
			"playlist": playlist,
		},
	))

}

func GetUserPlaylistsHandler(c *gin.Context) {

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	playlists, err := GetUserPlaylistsService(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"user playlists fetched successfully",
		gin.H{
			"playlist": playlists,
		},
	))
}

func GetUserPlaylistByIdHandler(c *gin.Context) {

	playlistId, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
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

	playlist, err := GetUserPlaylistByIdService(playlistId, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"user playlist fetched successfully",
		gin.H{
			"playlist": playlist,
		},
	))

}

func GetCoursePlaylistByIdHandler(c *gin.Context) {

	playlistId, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
			nil,
		))
		return
	}

	playlist, err := GetCoursePlaylistByIdService(playlistId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"course playlist fetched successfully",
		gin.H{
			"playlist": playlist,
		},
	))
}

func AddVideoToPlaylistHandler(c *gin.Context) {

	videoId, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	playlistId, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
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

	err = AddVideoToPlaylistService(playlistId, videoId, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"video added to playlist successfully",
		nil,
	))

}

func DeleteVideoFromPlaylistHandler(c *gin.Context) {

	videoId, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	playlistId, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
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

	err = DeleteVideoFromPlaylistService(playlistId, videoId, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"video deleted from playlist successfully",
		nil,
	))
}

func UpdatePlaylistHandler(c *gin.Context) {
	/**
	TODO:
		Take playlistId
		and update the playlist
	*/

	playlistId, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
			nil,
		))
		return
	}

	var req UpdatePlaylistRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
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

	playlist, err := UpdatePlaylistService(req, playlistId, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"playlist updated successfully",
		playlist,
	))
}
