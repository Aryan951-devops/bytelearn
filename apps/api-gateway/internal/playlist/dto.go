// Package playlist manages user custom video collections and course playlists.
package playlist

import (
	"time"

	"github.com/google/uuid"
)

// UpdatePlaylistRequest defines payload for updating a playlist.
type UpdatePlaylistRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// CreatePlaylistRequest defines payload for creating a playlist.
type CreatePlaylistRequest struct {
	Type        string     `json:"type" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	CourseID    *uuid.UUID `json:"course_id"`
}

// VideoMetadata holds snippet details of a video inside a playlist.
type VideoMetadata struct {
	ID           uuid.UUID `json:"video_id"`
	Title        string    `json:"title"`
	ThumbnailURL *string   `json:"thumbnail_url"`
	Views        int64     `json:"views"`
}

// PlaylistResponse represents the playlist output structure.
type PlaylistResponse struct {
	ID          uuid.UUID       `json:"playlist_id"`
	Type        string          `json:"type"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	UserID      uuid.UUID       `json:"user_id"`
	CourseID    *uuid.UUID      `json:"course_id"`
	Videos      []VideoMetadata `json:"videos"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

// CoursePlaylistResponse represents specialized course playlist output structure.
type CoursePlaylistResponse struct {
	PlaylistResponse
	EducatorID         string  `json:"educator_user_id"`
	EducatorUsername   string  `json:"educator_username"`
	EducatorName       string  `json:"educator_name"`
	EducatorProfilePic *string `json:"educator_profile_pic"`
}
