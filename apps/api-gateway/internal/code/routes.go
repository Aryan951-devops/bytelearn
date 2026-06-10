package code

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes defines the HTTP routes for the code package.
func RegisterRoutes(router *gin.RouterGroup) {

	code := router.Group("/code")

	code.POST("/cp/", auth.Middleware(), CreateCodingPracticeHandler)
	code.GET("/cp/:contestId", GetCodingPracticeHandler)
	code.GET("/cp/playlist/:playlistId", GetCodingPracticesOfPlaylistHandler)
	code.GET("/cp/question/:questionId", GetCodingQuestionHandler)
	code.POST("/cp/question", auth.Middleware(), CreateCodingQuestionHandler)
	code.POST("cp/testcase", auth.Middleware(), CreateCodingTestCaseHandler)
	code.GET("/cp/testcases-sample/:questionId", GetSampleTestCasesHandler)
	code.POST("/cp/submit", auth.Middleware(), SubmitCodeHandler)
	code.POST("/cp/submit-sample", auth.Middleware(), SubmitSampleCodeHandler)
	code.GET("/cp/poll/:submissionId", auth.Middleware(), GetSubmissionStatusHandler)
	code.GET("/cp/submission-result/:submissionId", auth.Middleware(), GetSubmissionResultHandler)

}
