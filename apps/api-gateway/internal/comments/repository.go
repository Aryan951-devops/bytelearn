package comments

import (
	"context"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// GetAllCommentOfVideo gets all comments for a specific video.
func GetAllCommentOfVideo(videoID uuid.UUID) ([]CommentResponse, error) {
	query := `
		SELECT c.comment_id, c.video_id, c.user_id, 
		u.username, c.content, c.created_at, c.updated_at
		FROM comments c INNER JOIN users u
		ON c.user_id = u.user_id
		WHERE video_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		videoID,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	comments := []CommentResponse{}

	for rows.Next() {
		var c CommentResponse

		err := rows.Scan(
			&c.ID,
			&c.VideoID,
			&c.UserID,
			&c.Username,
			&c.Content,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		comments = append(comments, c)
	}

	return comments, nil
}

// CreateComment saves a new comment to the database.
func CreateComment(comment *models.Comment) (*models.Comment, error) {
	query := `
		INSERT INTO comments 
		(video_id, user_id, content, created_at, updated_at)
		VALUES
		($1, $2, $3, NOW(), NOW())
		RETURNING
		comment_id, video_id, user_id, content,
		created_at, updated_at
	`

	var newComment models.Comment
	err := database.DB.QueryRow(
		context.Background(),
		query,
		comment.VideoID,
		comment.UserID,
		comment.Content,
	).Scan(
		&newComment.ID,
		&newComment.VideoID,
		&newComment.UserID,
		&newComment.Content,
		&newComment.CreatedAt,
		&newComment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newComment, nil
}

// DeleteComment removes a comment from the database by its ID.
func DeleteComment(commentID uuid.UUID) error {
	query := `
		DELETE FROM comments 
		WHERE comment_id = $1
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		commentID,
	)

	return err
}

// UpdateComment saves changes to an existing comment.
func UpdateComment(comment *models.Comment) (*models.Comment, error) {
	query := `
		UPDATE comments
		SET content =  $1
		WHERE comment_id = $2
		RETURNING
		comment_id, video_id, user_id, content,
		created_at, updated_at
	`

	var updatedComment models.Comment
	err := database.DB.QueryRow(
		context.Background(),
		query,
		comment.Content,
		comment.ID,
	).Scan(
		&updatedComment.ID,
		&updatedComment.VideoID,
		&updatedComment.UserID,
		&updatedComment.Content,
		&updatedComment.CreatedAt,
		&updatedComment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedComment, nil
}

// GetCommentByID finds a single comment by its ID.
func GetCommentByID(commentID uuid.UUID) (*models.Comment, error) {

	query := `
		SELECT
			comment_id,
			video_id,
			user_id,
			content,
			created_at,
			updated_at
		FROM comments
		WHERE comment_id = $1
	`

	var comment models.Comment

	err := database.DB.QueryRow(
		context.Background(),
		query,
		commentID,
	).Scan(
		&comment.ID,
		&comment.VideoID,
		&comment.UserID,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &comment, nil
}
