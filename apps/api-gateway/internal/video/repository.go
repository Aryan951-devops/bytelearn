package video

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
	"github.com/google/uuid"
)

func FetchAllVideosOfUser(user_id uuid.UUID) (*[]models.Video, error) {
	query := `
		SELECT video_id, title, description,
		videofile_url, videofile_public_id, thumbnail_url,
		thumbnail_public_id, duration_seconds,
		views, uploaded_by, created_at, updated_at
		FROM videos
		WHERE uploaded_by = $1
	`

	rows, err := database.DB.Query(context.Background(), query, user_id)
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

func FetchVideoByID(video_id uuid.UUID) (*models.Video, error) {
	query := `
		SELECT 
			video_id, title, description,
			videofile_url, videofile_public_id, thumbnail_url,
			thumbnail_public_id, duration_seconds,
			views, uploaded_by, created_at, updated_at
		FROM videos
		WHERE video_id = $1
	`

	var video models.Video

	err := database.DB.QueryRow(
		context.Background(),
		query,
		video_id,
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

	return &video, nil
}

func DeleteVideoByID(video_id uuid.UUID) error {
	query := `
		DELETE FROM videos
		WHERE video_id = $1
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		video_id,
	)

	return err
}

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

	var updated_video models.Video
	err := database.DB.QueryRow(context.Background(),
		query,
		video.Title,
		video.Description,
		video.Thumbnail_Url,
		video.Thumbnail_PublicID,
		video.ID,
	).Scan(
		&updated_video.ID,
		&updated_video.Title,
		&updated_video.Description,
		&updated_video.Videofile_Url,
		&updated_video.Videofile_PublicID,
		&updated_video.Thumbnail_Url,
		&updated_video.Thumbnail_PublicID,
		&updated_video.DurationSeconds,
		&updated_video.Views,
		&updated_video.UploadedBy,
		&updated_video.CreatedAt,
		&updated_video.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("video not found")
		}
		return nil, err
	}

	return &updated_video, nil
}
