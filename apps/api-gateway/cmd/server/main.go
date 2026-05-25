package main

import (
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/user"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/video"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/config"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
)

func main() {
	config.LoadConfig()

	database.ConnectDB(config.AppConfig.DatabaseURL)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)

	allowedOrigins := strings.Split(config.AppConfig.ALLOWED_ORIGINS, ",")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API Versioning
	v1 := router.Group("/api/v1")

	// Register Feature Routes
	auth.RegisterRoutes(v1)
	user.RegisterRoutes(v1)
	video.RegisterRoutes(v1)

	log.Printf("Server running on port %s", config.AppConfig.Port)

	if err := router.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
