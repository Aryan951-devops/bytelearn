package video

import "github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/models"

func GetAllVideos() (*[]models.Video, error) {
	videos, err := FetchAllVideos()

	if err != nil {
		return nil, err
	}

	return videos, nil
}
