package playlist

import (
	"errors"
	"strings"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreatePlaylistService processes payload validation to generate a playlist.
func CreatePlaylistService(req CreatePlaylistRequest,
	userID uuid.UUID) (*models.Playlist, error) {

	playlist := models.Playlist{
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		CourseID:    req.CourseID,
	}

	return CreatePlaylist(&playlist)
}

// GetUserPlaylistsService fetches playlists belonging to a user.
func GetUserPlaylistsService(userID uuid.UUID,
) ([]models.Playlist, error) {

	return GetUserPlaylists(userID)
}

// GetUserPlaylistByIDService fetches a singular user playlist entity.
func GetUserPlaylistByIDService(playlistID uuid.UUID,
	userID uuid.UUID,
) (*PlaylistResponse, error) {

	if err := VerifyPlaylistOwnership(playlistID, userID); err != nil {
		return nil, err
	}

	return GetUserPlaylistWithVideos(playlistID)
}

// GetCoursePlaylistByIDService retrieves a specific course playlist item.
func GetCoursePlaylistByIDService(playlistID uuid.UUID,
) (*CoursePlaylistResponse, error) {

	return GetCoursePlaylist(playlistID)
}

// AddVideoToPlaylistService appends a selected video metadata into a playlist.
func AddVideoToPlaylistService(playlistID uuid.UUID,
	videoID uuid.UUID,
	userID uuid.UUID,
) error {

	if err := VerifyPlaylistOwnership(playlistID, userID); err != nil {
		return err
	}

	return AddVideoToPlaylist(playlistID, videoID)
}

// DeleteVideoFromPlaylistService handles removal of videos from custom playlists.
func DeleteVideoFromPlaylistService(playlistID uuid.UUID,
	videoID uuid.UUID,
	userID uuid.UUID,
) error {

	if err := VerifyPlaylistOwnership(playlistID, userID); err != nil {
		return err
	}

	return DeleteVideoFromPlaylist(playlistID, videoID)
}

// UpdatePlaylistService updates meta settings of an existing playlist.
func UpdatePlaylistService(req UpdatePlaylistRequest,
	playlistID uuid.UUID,
	userID uuid.UUID,
) (*models.Playlist, error) {

	existingPlaylist, err := GetPlaylistByID(playlistID)
	if err != nil {
		return nil, err
	}

	if existingPlaylist.UserID != userID {
		return nil, errors.New("you are not authorized to update this playlist")
	}

	if req.Title != nil && strings.TrimSpace(*req.Title) != "" {
		existingPlaylist.Title = *req.Title
	}

	if req.Description != nil && strings.TrimSpace(*req.Description) != "" {
		existingPlaylist.Description = req.Description
	}

	return UpdatePlaylist(existingPlaylist)
}
