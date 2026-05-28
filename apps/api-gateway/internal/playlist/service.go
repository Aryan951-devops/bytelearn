package playlist

import (
	"errors"
	"strings"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
)

func CreatePlaylistService(req CreatePlaylistRequest,
	user_id uuid.UUID) (*models.Playlist, error) {

	playlist := models.Playlist{
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
		UserID:      user_id,
		CourseID:    req.CourseID,
	}

	return CreatePlaylist(&playlist)
}

func GetUserPlaylistsService(user_id uuid.UUID,
) ([]models.Playlist, error) {

	return GetUserPlaylists(user_id)
}

func GetUserPlaylistByIdService(playlist_id uuid.UUID,
	user_id uuid.UUID,
) (*PlaylistResponse, error) {

	if err := VerifyPlaylistOwnership(playlist_id, user_id); err != nil {
		return nil, err
	}

	return GetUserPlaylistWithVideos(playlist_id)
}

func GetCoursePlaylistByIdService(playlist_id uuid.UUID,
) (*CoursePlaylistResponse, error) {

	return GetCoursePlaylist(playlist_id)
}

func AddVideoToPlaylistService(playlist_id uuid.UUID,
	video_id uuid.UUID,
	user_id uuid.UUID,
) error {

	if err := VerifyPlaylistOwnership(playlist_id, user_id); err != nil {
		return err
	}

	return AddVideoToPlaylist(playlist_id, video_id)
}

func DeleteVideoFromPlaylistService(playlist_id uuid.UUID,
	video_id uuid.UUID,
	user_id uuid.UUID,
) error {

	if err := VerifyPlaylistOwnership(playlist_id, user_id); err != nil {
		return err
	}

	return DeleteVideoFromPlaylist(playlist_id, video_id)
}

func UpdatePlaylistService(req UpdatePlaylistRequest,
	playlist_id uuid.UUID,
	user_id uuid.UUID,
) (*models.Playlist, error) {

	existingPlaylist, err := GetPlaylistById(playlist_id)
	if err != nil {
		return nil, err
	}

	if existingPlaylist.UserID != user_id {
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
