package playlist

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreatePlaylistHandler handles REST endpoints for creating playlist.
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

// GetUserPlaylistsHandler handles requests fetching multiple playlist for authenticated users.
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

// GetUserPlaylistByIDHandler parses route IDs to fetch user custom playlists.
func GetUserPlaylistByIDHandler(c *gin.Context) {

	playlistID, err := uuid.Parse(c.Param("playlistId"))
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

	playlist, err := GetUserPlaylistByIDService(playlistID, user.ID)
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

// GetCoursePlaylistByIDHandler handles request for fetching course playlist.
func GetCoursePlaylistByIDHandler(c *gin.Context) {

	playlistID, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
			nil,
		))
		return
	}

	playlist, err := GetCoursePlaylistByIDService(playlistID)
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

// AddVideoToPlaylistHandler handles requests to add video into playlist.
func AddVideoToPlaylistHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	playlistID, err := uuid.Parse(c.Param("playlistId"))
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

	err = AddVideoToPlaylistService(playlistID, videoID, user.ID)

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

// DeleteVideoFromPlaylistHandler deletes video inside a playlist.
func DeleteVideoFromPlaylistHandler(c *gin.Context) {

	videoID, err := uuid.Parse(c.Param("videoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid video ID format",
			nil,
		))
		return
	}

	playlistID, err := uuid.Parse(c.Param("playlistId"))
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

	err = DeleteVideoFromPlaylistService(playlistID, videoID, user.ID)

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

// UpdatePlaylistHandler handles updating of playlist.
func UpdatePlaylistHandler(c *gin.Context) {

	playlistID, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist ID format",
			nil,
		))
		return
	}

	var req UpdatePlaylistRequest

	if err = c.ShouldBindJSON(&req); err != nil {
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

	playlist, err := UpdatePlaylistService(req, playlistID, user.ID)

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
