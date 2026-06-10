// Package video manages video files data structures.
package video

// UploadVideoRequest holds entry fields to save a video.
type UploadVideoRequest struct {
	Title             string `json:"title" binding:"required"`
	Description       string `json:"description"`
	VideofileURL      string `json:"videofile_url" binding:"required"`
	VideofilePublicID string `json:"videofile_public_id" binding:"required"`
	DurationSeconds   int32  `json:"duration_seconds"`
}

// UpdateVideoRequest holds fields to change video metadata.
type UpdateVideoRequest struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}
