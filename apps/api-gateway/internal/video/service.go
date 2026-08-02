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

// GetAllVideos gets a list of all videos.
func GetAllVideos() (*[]models.Video, error) {
	videos, err := FetchAllVideos()

	if err != nil {
		return nil, err
	}

	return videos, nil
}

// GetAllVideosOfUserService gets a list of videos belonging to a user.
func GetAllVideosOfUserService(userID uuid.UUID) (*[]models.Video, error) {
	videos, err := FetchAllVideosOfUser(userID)

	if err != nil {
		return nil, err
	}

	return videos, nil
}

// GetVideo reads a single video profile.
func GetVideo(videoID uuid.UUID, userID uuid.UUID) (*models.Video, error) {
	video, err := FetchVideoByID(videoID, userID)

	if err != nil {
		return nil, err
	}

	return video, nil
}

// DeleteVideo removes a video file record.
func DeleteVideo(videoID uuid.UUID, userID uuid.UUID) (*models.Video, error) {
	video, err := FetchVideoByID(videoID, uuid.Nil)

	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	if video.UploadedBy != userID {
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

	err = DeleteVideoByID(videoID)
	if err != nil {
		return nil, err
	}

	return video, nil
}

// GenerateUploadSignature creates a signed key for media upload servers.
func GenerateUploadSignature() (gin.H, error) {
	timestamp := time.Now().Unix()

	paramsToSign := url.Values{}
	paramsToSign.Set("timestamp", strconv.FormatInt(timestamp, 10))
	paramsToSign.Set("folder", "bytelearn/videos")
	// paramsToSign.Set("resource_type", "video")

	signature, err := config.AppConfig.MediaUploader.SignParameters(
		paramsToSign,
		config.AppConfig.Cloudinary.CloudinaryAPISecret,
	)

	if err != nil {
		return nil, err
	}

	signatureData := gin.H{
		"timestamp":  timestamp,
		"signature":  signature,
		"api_key":    config.AppConfig.Cloudinary.CloudinaryAPIKey,
		"cloud_name": config.AppConfig.Cloudinary.CloudinaryCloudName,
		"folder":     "bytelearn/videos",
	}
	return signatureData, nil
}

// UploadVideo processes requests to add a video file.
func UploadVideo(req UploadVideoRequest, userID uuid.UUID,
) (*models.Video, error) {

	video := &models.Video{
		Title:              req.Title,
		Description:        &req.Description,
		Videofile_Url:      req.VideofileURL,
		Videofile_PublicID: req.VideofilePublicID,
		DurationSeconds:    req.DurationSeconds,
		UploadedBy:         userID,
	}

	return CreateVideo(video)
}

// UpdateVideo processes requests to edit video descriptions.
func UpdateVideo(req UpdateVideoRequest, userID uuid.UUID,
	videoID uuid.UUID, tempPath *string) (*models.Video, error) {

	video, err := GetVideo(videoID, uuid.Nil)
	if err != nil {
		return nil, err
	}

	if video.UploadedBy != userID {
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

	updatedVideo, err := UpdateVideoByID(video)

	if err != nil {
		return nil, err
	}

	return updatedVideo, nil
}
