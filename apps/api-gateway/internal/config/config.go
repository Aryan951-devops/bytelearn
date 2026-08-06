// Package config handles environment variable configurations.
package config

import (
	"log"
	"os"

	mediauploader "github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/mediaUploader"
	cloudinaryuploader "github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/mediaUploader/cloudinaryUploader"
	"github.com/Aryan951-devops/bytelearn/pkg/redis"
	"github.com/joho/godotenv"
)

// CloudinaryConfig holds API secret combinations for media servers.
type CloudinaryConfig struct {
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
}

// Config manages overall application system variables.
type Config struct {
	RedisClient       *redis.Client
	Port              string
	DatabaseURL       string
	RecommendationURL string
	JWTSecret         string
	AllowedOrigins    string
	RedisAddr         string
	Cloudinary        CloudinaryConfig
	MediaUploader     mediauploader.MediaUploader
}

// AppConfig stores the globally accessible configurations.
var AppConfig Config

// LoadConfig initializes and hydrates configuration elements.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	AppConfig = Config{
		Port:              getEnv("PORT", "8080"),
		DatabaseURL:       getEnv("DATABASE_URL", ""),
		RecommendationURL: getEnv("RECOMMENDATION_URL", "localhost:8001"),
		JWTSecret:         getEnv("JWT_SECRET", ""),
		AllowedOrigins:    getEnv("ALLOWED_ORIGINS", ""),
		RedisAddr:         getEnv("REDIS_ADDR", "localhost:6379"),
		Cloudinary: CloudinaryConfig{
			CloudinaryCloudName: getEnv("CLOUDINARY_CLOUD_NAME", ""),
			CloudinaryAPIKey:    getEnv("CLOUDINARY_API_KEY", ""),
			CloudinaryAPISecret: getEnv("CLOUDINARY_API_SECRET", ""),
		},
	}

	cloudinaryUploaderService, err := cloudinaryuploader.NewCloudinaryUploader(
		AppConfig.Cloudinary.CloudinaryAPIKey,
		AppConfig.Cloudinary.CloudinaryAPISecret,
		AppConfig.Cloudinary.CloudinaryCloudName,
	)

	if err != nil {
		log.Fatal("failed to initialize cloudinary uploader: ", err)
	}

	AppConfig.MediaUploader = cloudinaryUploaderService

	AppConfig.RedisClient = redis.MustConnectRedis(
		AppConfig.RedisAddr,
		"",
		0,
	)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
