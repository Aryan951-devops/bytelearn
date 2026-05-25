package video

import (
	"context"
	"errors"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"
)

func FetchAllVideos() (*[]models.Video, error) {
	query := `
		SELECT video_id, title, description,
		videofile_url, thumbnail_url, duration_seconds,
		views, uploaded_by, created_at, updated_at
		FROM videos
	`

	rows, err := database.DB.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var videos []models.Video

	for rows.Next() {
		var v models.Video

		err := rows.Scan(
			&v.ID,
			&v.Title,
			&v.Description,
			&v.VideofileUrl,
			&v.ThumbnailUrl,
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
