package video

import "time"

type VideoResponse struct {
	ID              string    `json:"video_id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ThumbnailURL    string    `json:"thumbnail_url"`
	VideoURL        string    `json:"videofile_url"`
	DurationSeconds string    `json:"duration_seconds"`
	Views           int32     `json:"views"`
	UploadedBy      string    `json:"uploaded_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
