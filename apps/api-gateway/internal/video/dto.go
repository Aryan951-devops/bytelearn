// Package video manages video files data structures.
package video

import "github.com/google/uuid"

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

// VideoMetadata holds snippet details of a video.
type VideoMetadata struct {
	ID            uuid.UUID `json:"video_id"`
	Title         string    `json:"title"`
	Thumbnail_Url *string   `json:"thumbnail_url"`
	Views         int64     `json:"views"`
}

// RecommendationSearchItem maps to the item returned by Python Recommendation Service
type RecommendationSearchItem struct {
	VideoID    uuid.UUID `json:"video_id"`
	Title      string    `json:"title"`
	Document   string    `json:"document"`
	Similarity float64   `json:"similarity"`
}

// RecommendationSearchResponse maps to the top-level response from Python
type RecommendationSearchResponse struct {
	Results []RecommendationSearchItem `json:"results"`
	Count   int                        `json:"count"`
}
