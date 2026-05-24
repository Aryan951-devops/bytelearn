package config

import (
	"log"
	"os"

	mediauploader "github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/mediaUploader"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/mediaUploader/cloudinaryUploader"
	"github.com/joho/godotenv"
)

type CloudinaryConfig struct {
	CLOUDINARY_CLOUD_NAME string
	CLOUDINARY_API_KEY    string
	CLOUDINARY_API_SECRET string
}

type Config struct {
	Port          string
	DatabaseURL   string
	JWT_SECRET    string
	Cloudinary    CloudinaryConfig
	MediaUploader mediauploader.MediaUploader
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	AppConfig = Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWT_SECRET:  getEnv("JWT_SECRET", ""),
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

}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
