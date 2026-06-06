package code

import (
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {

	code := router.Group("/code")

	code.POST("/cp/", auth.AuthMiddleware(), CreateCodingPracticeHandler)
	code.GET("/cp/:contestId", GetCodingPracticeHandler)
	code.POST("/cp/question", auth.AuthMiddleware(), CreateCodingQuestionHandler)
	code.POST("cp/testcase", auth.AuthMiddleware(), CreateCodingTestCaseHandler)
	code.GET("/cp/testcases-sample/:questionId", GetSampleTestCasesHandler)
	code.POST("/cp/submit", auth.AuthMiddleware(), SubmitCodeHandler)
	code.POST("/cp/submit-sample", auth.AuthMiddleware(), SubmitSampleCodeHandler)
	code.GET("/cp/poll/:submissionId", auth.AuthMiddleware(), GetSubmissionStatusHandler)
	code.GET("/cp/submission-result/:submissionId", auth.AuthMiddleware(), GetSubmissionResultHandler)

}
