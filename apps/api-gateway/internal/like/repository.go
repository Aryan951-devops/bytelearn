package like

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/google/uuid"
)

// IsVideoLiked checks if a user has liked a video.
func IsVideoLiked(userID, videoID uuid.UUID) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM video_likes
			WHERE user_id = $1
			AND video_id = $2
		)
	`

	var liked bool

	err := database.DB.QueryRow(
		context.Background(),
		query,
		userID,
		videoID,
	).Scan(&liked)

	return liked, err
}

// LikeVideo registers a like from a user for a video.
func LikeVideo(userID, videoID uuid.UUID) error {

	query := `
		INSERT INTO video_likes
		(user_id, video_id)
		VALUES ($1, $2)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		userID,
		videoID,
	)

	return err
}

// UnlikeVideo removes a user's like from a video.
func UnlikeVideo(userID, videoID uuid.UUID) error {

	query := `
		DELETE FROM video_likes
		WHERE user_id = $1
		AND video_id = $2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		userID,
		videoID,
	)

	return err
}

// GetTotalVideoLikes returns total likes for a video.
func GetTotalVideoLikes(videoID uuid.UUID) (int64, error) {

	query := `
		SELECT COUNT(*)
		FROM video_likes
		WHERE video_id = $1
	`

	var total int64

	err := database.DB.QueryRow(
		context.Background(),
		query,
		videoID,
	).Scan(&total)

	return total, err
}

// IsCommentLiked checks if a user has liked a comment.
func IsCommentLiked(userID, commentID uuid.UUID) (bool, error) {

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM comment_likes
			WHERE user_id = $1
			AND comment_id = $2
		)
	`

	var liked bool

	err := database.DB.QueryRow(
		context.Background(),
		query,
		userID,
		commentID,
	).Scan(&liked)

	return liked, err
}

// LikeComment registers a like from a user for a comment.
func LikeComment(userID, commentID uuid.UUID) error {

	query := `
		INSERT INTO comment_likes
		(user_id, comment_id)
		VALUES ($1, $2)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		userID,
		commentID,
	)

	return err
}

// UnlikeComment removes a user's like from a comment.
func UnlikeComment(userID, commentID uuid.UUID) error {

	query := `
		DELETE FROM comment_likes
		WHERE user_id = $1
		AND comment_id = $2
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		userID,
		commentID,
	)

	return err
}

// GetTotalCommentLikes returns total likes for a comment.
func GetTotalCommentLikes(commentID uuid.UUID) (int64, error) {

	query := `
		SELECT COUNT(*)
		FROM comment_likes
		WHERE comment_id = $1
	`

	var total int64

	err := database.DB.QueryRow(
		context.Background(),
		query,
		commentID,
	).Scan(&total)

	return total, err
}
