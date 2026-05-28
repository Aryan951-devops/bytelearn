package playlist

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	playlist := router.Group("/playlist")

	playlist.POST("/", auth.AuthMiddleware(), CreatePlaylistHandler)
	playlist.GET("/user-playlists", auth.AuthMiddleware(), GetUserPlaylistsHandler)
	playlist.PATCH("/add/:videoId/:playlistId", auth.AuthMiddleware(), AddVideoToPlaylistHandler)
	playlist.PATCH("/remove/:videoId/:playlistId", auth.AuthMiddleware(), DeleteVideoFromPlaylistHandler)
	playlist.PATCH("/:playlistId", auth.AuthMiddleware(), UpdatePlaylistHandler)
	playlist.GET("/course/:playlistId", GetCoursePlaylistByIdHandler)
	playlist.GET("/user/:playlistId", auth.AuthMiddleware(), GetUserPlaylistByIdHandler)

}
