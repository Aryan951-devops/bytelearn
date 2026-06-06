package config

import (
	"log"
	"os"

	mediauploader "github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/mediaUploader"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/mediaUploader/cloudinaryUploader"
	"github.com/Aryan951-devops/bytelearn/pkg/redis"
	"github.com/joho/godotenv"
)

type CloudinaryConfig struct {
	CLOUDINARY_CLOUD_NAME string
	CLOUDINARY_API_KEY    string
	CLOUDINARY_API_SECRET string
}

type Config struct {
	Port            string
	DatabaseURL     string
	JWT_SECRET      string
	ALLOWED_ORIGINS string
	RedisAddr       string
	Cloudinary      CloudinaryConfig
	MediaUploader   mediauploader.MediaUploader
	RedisClient     *redis.Client
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	AppConfig = Config{
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		JWT_SECRET:      getEnv("JWT_SECRET", ""),
		ALLOWED_ORIGINS: getEnv("ALLOWED_ORIGINS", ""),
		RedisAddr:       getEnv("REDIS_ADDR", "localhost:6789"),
		Cloudinary: CloudinaryConfig{
			CLOUDINARY_CLOUD_NAME: getEnv("CLOUDINARY_CLOUD_NAME", ""),
			CLOUDINARY_API_KEY:    getEnv("CLOUDINARY_API_KEY", ""),
			CLOUDINARY_API_SECRET: getEnv("CLOUDINARY_API_SECRET", ""),
		},
	}

	cloudinaryUploaderService, err := cloudinaryUploader.NewCloudinaryUploader(
		AppConfig.Cloudinary.CLOUDINARY_API_KEY,
		AppConfig.Cloudinary.CLOUDINARY_API_SECRET,
		AppConfig.Cloudinary.CLOUDINARY_CLOUD_NAME,
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
