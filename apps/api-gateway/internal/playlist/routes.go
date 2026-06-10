package playlist

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers endpoints for handling playlists.
func RegisterRoutes(router *gin.RouterGroup) {

	playlist := router.Group("/playlist")

	playlist.POST("/", auth.Middleware(), CreatePlaylistHandler)
	playlist.GET("/user-playlists", auth.Middleware(), GetUserPlaylistsHandler)
	playlist.PATCH("/add/:videoId/:playlistId", auth.Middleware(), AddVideoToPlaylistHandler)
	playlist.PATCH("/remove/:videoId/:playlistId", auth.Middleware(),
		DeleteVideoFromPlaylistHandler)
	playlist.PATCH("/:playlistId", auth.Middleware(), UpdatePlaylistHandler)
	playlist.GET("/course/:playlistId", GetCoursePlaylistByIDHandler)
	playlist.GET("/user/:playlistId", auth.Middleware(), GetUserPlaylistByIDHandler)

}
