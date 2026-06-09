package quiz

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	quiz := router.Group("/quiz")

	quiz.POST("/", auth.AuthMiddleware(), CreateQuizHandler)
	quiz.GET("/start/:quizId", auth.AuthMiddleware(), StartQuizHandler)
	quiz.POST("/submit/:attemptId", auth.AuthMiddleware(), SubmitQuizHandler)
	quiz.GET("/result/:attemptId", auth.AuthMiddleware(), GetAttemptResultHandler)
	quiz.GET("/attempts/:quizId", auth.AuthMiddleware(), GetAllAttemptsOfQuizHandler)
	quiz.GET("/quizzes/:playlistId", GetAllQuizesOfPlaylistHandler)
}
