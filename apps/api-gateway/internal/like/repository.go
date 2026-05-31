package like

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
	"github.com/google/uuid"
)

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
