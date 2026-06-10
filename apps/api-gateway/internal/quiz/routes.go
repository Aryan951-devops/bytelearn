package quiz

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the quiz package.
func RegisterRoutes(router *gin.RouterGroup) {

	quiz := router.Group("/quiz")

	quiz.POST("/", auth.Middleware(), CreateQuizHandler)
	quiz.GET("/start/:quizId", auth.Middleware(), StartQuizHandler)
	quiz.POST("/submit/:attemptId", auth.Middleware(), SubmitQuizHandler)
	quiz.GET("/result/:attemptId", auth.Middleware(), GetAttemptResultHandler)
	quiz.GET("/attempts/:quizId", auth.Middleware(), GetAllAttemptsOfQuizHandler)
	quiz.GET("/quizzes/:playlistId", GetAllQuizesOfPlaylistHandler)
}
