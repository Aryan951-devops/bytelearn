package video

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// FetchAllVideosOfUser finds all videos owned by a user.
func FetchAllVideosOfUser(userID uuid.UUID) (*[]models.Video, error) {
	query := `
		SELECT video_id, title, description,
		videofile_url, videofile_public_id, thumbnail_url,
		thumbnail_public_id, duration_seconds,
		views, uploaded_by, created_at, updated_at
		FROM videos
		WHERE uploaded_by = $1
	`

	rows, err := database.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	videos := []models.Video{}

	for rows.Next() {
		var v models.Video

		err := rows.Scan(
			&v.ID,
			&v.Title,
			&v.Description,
			&v.Videofile_Url,
			&v.Videofile_PublicID,
			&v.Thumbnail_Url,
			&v.Thumbnail_PublicID,
			&v.DurationSeconds,
			&v.Views,
			&v.UploadedBy,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		videos = append(videos, v)
	}

	return &videos, nil
}

// FetchAllVideos finds all recorded videos.
func FetchAllVideos() (*[]models.Video, error) {
	query := `
		SELECT video_id, title, description,
		videofile_url, videofile_public_id, thumbnail_url,
		thumbnail_public_id, duration_seconds,
		views, uploaded_by, created_at, updated_at
		FROM videos
	`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	videos := []models.Video{}

	for rows.Next() {
		var v models.Video

		err := rows.Scan(
			&v.ID,
			&v.Title,
			&v.Description,
			&v.Videofile_Url,
			&v.Videofile_PublicID,
			&v.Thumbnail_Url,
			&v.Thumbnail_PublicID,
			&v.DurationSeconds,
			&v.Views,
			&v.UploadedBy,
			&v.CreatedAt,
			&v.UpdatedAt,
		)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		videos = append(videos, v)
	}

	return &videos, nil
}

// FetchVideoByID finds a single video using its ID.
func FetchVideoByID(videoID uuid.UUID, userID uuid.UUID) (*models.Video, error) {
	query := `
		UPDATE videos
			SET views = views + 1,
			updated_at = NOW()
		WHERE video_id = $1
		RETURNING 
			video_id, title, description,
			videofile_url, videofile_public_id, thumbnail_url,
			thumbnail_public_id, duration_seconds,
			views, uploaded_by, created_at, updated_at
	`

	var video models.Video

	err := database.DB.QueryRow(
		context.Background(),
		query,
		videoID,
	).Scan(
		&video.ID,
		&video.Title,
		&video.Description,
		&video.Videofile_Url,
		&video.Videofile_PublicID,
		&video.Thumbnail_Url,
		&video.Thumbnail_PublicID,
		&video.DurationSeconds,
		&video.Views,
		&video.UploadedBy,
		&video.CreatedAt,
		&video.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	log.Println("UserID: ", userID)

	if userID != uuid.Nil {
		watchHistoryQuery := `
			INSERT INTO watch_history 
				(user_id, video_id, resume_time, updated_at)
			VALUES ($1, $2, 0, NOW())
			ON CONFLICT (user_id, video_id) 
			DO UPDATE SET 
				updated_at = NOW();
		`

		_, err := database.DB.Exec(
			context.Background(),
			watchHistoryQuery,
			userID,
			videoID,
		)

		if err != nil {
			// Log the error and still return the video so viewing isn't blocked by a history error
			log.Printf("failed to update watch history for user %s: %v", userID, err)
		}
	}

	return &video, nil
}

// DeleteVideoByID removes a video entry using its ID.
func DeleteVideoByID(videoID uuid.UUID) error {
	query := `
		DELETE FROM videos
		WHERE video_id = $1
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		videoID,
	)

	return err
}

// CreateVideo saves a new video record.
func CreateVideo(video *models.Video) (*models.Video, error) {

	query := `
		INSERT INTO videos (
			title, description, videofile_url,
			videofile_public_id, uploaded_by
		)
		VALUES ($1,$2,$3,$4,$5)
		RETURNING
			video_id, title, description,
			videofile_url, videofile_public_id,
			thumbnail_url, thumbnail_public_id,
			duration_seconds, views, uploaded_by,
			created_at, updated_at
	`

	var createdVideo models.Video

	err := database.DB.QueryRow(
		context.Background(),
		query,
		video.Title,
		video.Description,
		video.Videofile_Url,
		video.Videofile_PublicID,
		video.UploadedBy,
	).Scan(
		&createdVideo.ID,
		&createdVideo.Title,
		&createdVideo.Description,
		&createdVideo.Videofile_Url,
		&createdVideo.Videofile_PublicID,
		&createdVideo.Thumbnail_Url,
		&createdVideo.Thumbnail_PublicID,
		&createdVideo.DurationSeconds,
		&createdVideo.Views,
		&createdVideo.UploadedBy,
		&createdVideo.CreatedAt,
		&createdVideo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdVideo, nil
}

// UpdateVideoByID saves edits made to a video file.
func UpdateVideoByID(video *models.Video) (*models.Video, error) {

	query := `
		UPDATE videos
		SET
			title = $1,
			description = COALESCE($2, description),
			thumbnail_url = COALESCE($3, thumbnail_url),
			thumbnail_public_id = COALESCE($4, thumbnail_public_id)
		WHERE video_id = $5
		RETURNING
			video_id, title, description,
			videofile_url, videofile_public_id, thumbnail_url,
			thumbnail_public_id, duration_seconds,
			views, uploaded_by, created_at, updated_at
	`

	var updatedVideo models.Video
	err := database.DB.QueryRow(context.Background(),
		query,
		video.Title,
		video.Description,
		video.Thumbnail_Url,
		video.Thumbnail_PublicID,
		video.ID,
	).Scan(
		&updatedVideo.ID,
		&updatedVideo.Title,
		&updatedVideo.Description,
		&updatedVideo.Videofile_Url,
		&updatedVideo.Videofile_PublicID,
		&updatedVideo.Thumbnail_Url,
		&updatedVideo.Thumbnail_PublicID,
		&updatedVideo.DurationSeconds,
		&updatedVideo.Views,
		&updatedVideo.UploadedBy,
		&updatedVideo.CreatedAt,
		&updatedVideo.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("video not found")
		}
		return nil, err
	}

	return &updatedVideo, nil
}
