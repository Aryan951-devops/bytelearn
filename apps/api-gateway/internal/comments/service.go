package comments

import (
	"errors"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

func GetAllCommentOfVideoService(videoID uuid.UUID) ([]CommentResponse, error) {
	return GetAllCommentOfVideo(videoID)
}

func CreateCommentService(req CommentRequest,
	userID uuid.UUID,
	videoID uuid.UUID) (*models.Comment, error) {

	comment := models.Comment{
		VideoID: videoID,
		UserID:  userID,
		Content: req.Content,
	}
	return CreateComment(&comment)
}

func DeleteCommentService(
	commentID uuid.UUID,
	userID uuid.UUID,
) error {

	comment, err := GetCommentByID(commentID)
	if err != nil {
		return err
	}

	if comment.UserID != userID {
		return errors.New("you are not allowed to delete this comment")
	}

	return DeleteComment(commentID)
}

func UpdateCommentService(
	req CommentRequest,
	userID uuid.UUID,
	commentID uuid.UUID,
) (*models.Comment, error) {

	existingComment, err := GetCommentByID(commentID)
	if err != nil {
		return nil, err
	}

	if existingComment.UserID != userID {
		return nil, errors.New("you are not allowed to update this comment")
	}

	existingComment.Content = req.Content

	return UpdateComment(existingComment)
}
