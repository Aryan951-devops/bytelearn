package course

import (
	"time"

	"github.com/google/uuid"
)

type CreateCourseRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
}

type UpdateCourseRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Category    *string `json:"category"`
}

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

type EducatorDetail struct {
	ID             uuid.UUID `json:"user_id"`
	Username       string    `json:"username"`
	Name           string    `json:"name"`
	ProfilePic_Url string    `json:"profile_pic_url"`
}

type PlaylistPreview struct {
	ID          uuid.UUID `json:"playlist_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
}
