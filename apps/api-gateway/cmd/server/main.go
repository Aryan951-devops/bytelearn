package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "ByteLearn API running",
		})
	})

	log.Println("Server running on port 8080")

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
