package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/config"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/pkg/database"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDB(cfg.DatabaseURL)

	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "ByteLearn API running",
		})
	})

	log.Printf("Server running on port %s", cfg.Port)

	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
