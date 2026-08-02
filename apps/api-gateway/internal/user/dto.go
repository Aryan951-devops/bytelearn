// Package user handles user profile data and adjustments.
package user

import (
	"github.com/google/uuid"
)

// ChangePasswordRequest holds data to change a password.
type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// UpdateAccountRequest holds data to update account settings.
type UpdateAccountRequest struct {
	PhoneNo string `form:"phone_no"`
	City    string `form:"city"`
	State   string `form:"state"`
	Pincode string `form:"pincode"`
}

// VideoMetadata holds snippet details of a video.
type VideoMetadata struct {
	ID           uuid.UUID `json:"video_id"`
	Title        string    `json:"title"`
	ThumbnailURL *string   `json:"thumbnail_url"`
	Views        int64     `json:"views"`
}

// HistoryResponse returns the history output structure.
type HistoryResponse struct {
	UserID uuid.UUID       `json:"user_id"`
	Videos []VideoMetadata `json:"videos"`
}
