// Package course manages layout architectures for curated educational series.
package course

import (
	"time"

	"github.com/google/uuid"
)

// CreateCourseRequest accepts metadata values for initializing a course map.
type CreateCourseRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
}

// UpdateCourseRequest updates configurable values of target courses.
type UpdateCourseRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
}

// CourseWithPlaylistResponse embeds structural playlist information within a course profile.
type CourseWithPlaylistResponse struct {
	ID          uuid.UUID         `json:"course_id"`
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Category    *string           `json:"category"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	Playlists   []PlaylistPreview `json:"playlists"`
	Educators   []EducatorDetail  `json:"educators"`
}

// EducatorDetail shapes individual details of an instructional author.
type EducatorDetail struct {
	ID            uuid.UUID `json:"user_id"`
	Username      string    `json:"username"`
	Name          string    `json:"name"`
	ProfilePicURL *string   `json:"profile_pic_url"`
}

// PlaylistPreview provides a concise snapshot of a playlist's properties.
type PlaylistPreview struct {
	ID          uuid.UUID `json:"playlist_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
}
