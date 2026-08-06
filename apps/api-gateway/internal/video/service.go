package video

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
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

	res, err := CreateVideo(video)

	if err != nil {
		return nil, err
	}

	job := models.RecommendationJob{
		VidoeID:   res.ID,
		EventType: "video_created",
	}

	err = config.AppConfig.RedisClient.Publish(
		context.Background(),
		"recommendation_jobs",
		job,
	)

	if err != nil {
		fmt.Println("Got Error while publishing job to Redis")
	}

	return res, err
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

var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// CreateNewHttpRequest creates a new HTTP request with the searchText.
func CreateNewHttpRequest(searchText string,
) (*http.Request, error) {

	reqURL, err := url.Parse(fmt.Sprintf("%s/search",
		config.AppConfig.RecommendationURL))
	if err != nil {
		return nil, fmt.Errorf("invalid recommendation service url: %w", err)
	}

	query := reqURL.Query()
	query.Set("q", searchText)
	query.Set("limit", "10")
	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(context.Background(),
		http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return req, nil
}

// Returns the recommended videos while searching.
func GetVideosBySearchService(searchText string,
) (*[]VideoMetadata, error) {

	if searchText == "" {
		return &[]VideoMetadata{}, nil
	}

	req, err := CreateNewHttpRequest(searchText)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling recommendation service: %w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("recommendation service returned status code %d", resp.StatusCode)
	}

	var recResp RecommendationSearchResponse
	if err = json.NewDecoder(resp.Body).Decode(&recResp); err != nil {
		return nil, fmt.Errorf("failed to decode recommendation response: %w", err)
	}

	if len(recResp.Results) == 0 {
		return &[]VideoMetadata{}, nil
	}

	fmt.Println("Hey Response: ", recResp)
	videoIDs := make([]uuid.UUID, 0, recResp.Count)
	orderMap := make(map[uuid.UUID]int, recResp.Count)

	for index, item := range recResp.Results {
		videoIDs = append(videoIDs, item.VideoID)
		orderMap[item.VideoID] = index
	}

	fmt.Println("videoIDs: ", videoIDs)
	videos, err := FetchVideoMetaData(videoIDs)
	if err != nil {
		fmt.Println("Err: ", err)
		return nil, err
	}

	fetchedMap := make(map[uuid.UUID]VideoMetadata, len(videoIDs))
	for _, video := range *videos {
		fetchedMap[video.ID] = video
	}

	orderedVideos := make([]VideoMetadata, 0, len(fetchedMap))
	for _, id := range videoIDs {
		if video, exists := fetchedMap[id]; exists {
			orderedVideos = append(orderedVideos, video)
		}
	}

	return &orderedVideos, nil
}
