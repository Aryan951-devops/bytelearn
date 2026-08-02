package auth

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// OptionalMiddleware attempts to validate the user but allows anonymous access if unauthenticated.
func OptionalMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.Next()
			return
		}

		token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.Next()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Next()
			return
		}

		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			c.Next()
			return
		}

		parsedUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.Next()
			return
		}

		user, err := GetUserByUserID(parsedUUID)
		if err != nil {
			c.Next()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// Middleware protects routes by validating tokens.
func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString, err := c.Cookie("access_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfig.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			c.Abort()
			return
		}

		userIDStr, _ := claims["user_id"].(string)

		parsedUUID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "malformed user identifier format"})
			c.Abort()
			return
		}

		user, err := GetUserByUserID(parsedUUID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
