package comments

import "github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"

type CommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type CommentResponse struct {
	models.Comment
	Username string `json:"commented_by"`
}
