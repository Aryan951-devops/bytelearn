package video

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/config"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllVideos() (*[]models.Video, error) {
	videos, err := FetchAllVideos()

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func GetAllVideosOfUserService(user_id uuid.UUID) (*[]models.Video, error) {
	videos, err := FetchAllVideosOfUser(user_id)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

func GetVideo(video_id uuid.UUID) (*models.Video, error) {
	video, err := FetchVideoByID(video_id)

	if err != nil {
		return nil, err
	}

	return video, nil
}

func DeleteVideo(video_id uuid.UUID, user_id uuid.UUID) (*models.Video, error) {
	video, err := FetchVideoByID(video_id)

	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if video.UploadedBy != user_id {
		return nil, errors.New("you are not authorized to delete this video")
	}

	err = config.AppConfig.MediaUploader.DeleteByPublicID(video.Videofile_PublicID)
	if err != nil {
		log.Println("failed to delete video on cloudinary:", err)
	}

	err = config.AppConfig.MediaUploader.DeleteByPublicID(*video.Thumbnail_PublicID)
	if err != nil {
		log.Println("failed to delete thumbnail on cloudinary:", err)
	}

	err = DeleteVideoByID(video_id)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func GenerateUploadSignature() (gin.H, error) {
	timestamp := time.Now().Unix()

	paramsToSign := url.Values{}
	paramsToSign.Set("timestamp", strconv.FormatInt(timestamp, 10))
	paramsToSign.Set("folder", "bytelearn/videos")
	// paramsToSign.Set("resource_type", "video")

	signature, err := config.AppConfig.MediaUploader.SignParameters(
		paramsToSign,
		config.AppConfig.Cloudinary.CLOUDINARY_API_SECRET,
	)

	if err != nil {
		return nil, err
	}

	signatureData := gin.H{
		"timestamp":  timestamp,
		"signature":  signature,
		"api_key":    config.AppConfig.Cloudinary.CLOUDINARY_API_KEY,
		"cloud_name": config.AppConfig.Cloudinary.CLOUDINARY_CLOUD_NAME,
		"folder":     "bytelearn/videos",
	}
	return signatureData, nil
}

func UploadVideo(req UploadVideoRequest, userID uuid.UUID,
) (*models.Video, error) {

	video := &models.Video{
		Title:              req.Title,
		Description:        &req.Description,
		Videofile_Url:      req.Videofile_Url,
		Videofile_PublicID: req.Videofile_PublicID,
		DurationSeconds:    req.DurationSeconds,
		UploadedBy:         userID,
	}

	return CreateVideo(video)
}

func UpdateVideo(req UpdateVideoRequest, user_id uuid.UUID,
	video_id uuid.UUID, tempPath *string) (*models.Video, error) {

	video, err := GetVideo(video_id)
	if err != nil {
		return nil, err
	}

	if video.UploadedBy != user_id {
		return nil, errors.New("you are not unauthorized to update video")
	}

	if strings.TrimSpace(req.Title) != "" {
		video.Title = req.Title
	}

	if strings.TrimSpace(req.Description) != "" {
		video.Description = &req.Description
	}

	if tempPath != nil {
		imageData := config.AppConfig.MediaUploader.UploadThumbnail(*tempPath)

		if imageData.Err != nil {
			return nil, imageData.Err
		}

		video.Thumbnail_Url = &imageData.PublicURL
		video.Thumbnail_PublicID = &imageData.PublicID
	}

	updated_video, err := UpdateVideoByID(video)

	if err != nil {
		return nil, err
	}

	return updated_video, nil
}
