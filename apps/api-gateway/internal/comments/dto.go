// Package comments handles the data and logic for user comments.
package comments

import "github.com/Aryan951-devops/bytelearn/pkg/models"

// CommentRequest holds the data coming in when creating or updating a comment.
type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}

// CommentResponse holds the data sent out when displaying a comment.
type CommentResponse struct {
	models.Comment
	Username string `json:"commented_by"`
}
