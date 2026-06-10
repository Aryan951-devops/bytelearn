package code

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateCodingPracticeHandler handles HTTP tasks for making practice elements.
func CreateCodingPracticeHandler(c *gin.Context) {

	var req CreateCodingPracticeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if user.Role != "educator" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are not authorized to create contest",
			nil,
		))
		return
	}

	practice, err := CreateCodingPracticeService(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"coding practice created successfully",
		gin.H{
			"practice": practice,
		},
	))
}

// GetCodingPracticeHandler handles HTTP tasks to inspect a practice set.
func GetCodingPracticeHandler(c *gin.Context) {

	contestID, err := uuid.Parse(c.Param("contestId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid contest ID format",
			nil,
		))
		return
	}

	practice, err := GetCodingPracticeService(contestID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"coding practices fetched successfully",
		gin.H{
			"practice": practice,
		},
	))
}

// GetCodingPracticesOfPlaylistHandler displays all practices of playlist.
func GetCodingPracticesOfPlaylistHandler(c *gin.Context) {

	playlistID, err := uuid.Parse(c.Param("playlistId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist id format",
			nil,
		))
		return
	}

	cps, err := GetCodingPracticesOfPlaylistService(playlistID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"coding practices fetched successfully",
		gin.H{
			"practices": cps,
		},
	))
}

// GetCodingQuestionHandler handles HTTP requests to view specific problems.
func GetCodingQuestionHandler(c *gin.Context) {

	questionID, err := uuid.Parse(c.Param("questionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid question id format",
			nil,
		))
		return
	}

	question, err := GetCodingQuestionService(questionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"coding question fetched successfully",
		gin.H{
			"question": question,
		},
	))

}

// CreateCodingQuestionHandler processes creation of coding question.
func CreateCodingQuestionHandler(c *gin.Context) {

	var req CreateCodingQuestionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if user.Role != "educator" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are not authorized to create question",
			nil,
		))
		return
	}

	question, err := CreateCodingQuestionService(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"coding question created successfully",
		gin.H{
			"question": question,
		},
	))
}

// CreateCodingTestCaseHandler processes testcase attached to specific problems.
func CreateCodingTestCaseHandler(c *gin.Context) {

	var req CreateTestCaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if user.Role != "educator" {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"you are not authorized to create testcase",
			nil,
		))
		return
	}

	testcase, err := CreateCodingTestCaseService(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"testcase created successfully",
		gin.H{
			"testcase": testcase,
		},
	))
}

// GetSampleTestCasesHandler reads visible testing items.
func GetSampleTestCasesHandler(c *gin.Context) {

	questionID, err := uuid.Parse(c.Param("questionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid question id format",
			nil,
		))
		return
	}

	testcases, err := GetSampleTestCasesService(questionID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"sample testcases fetched successfully",
		gin.H{
			"testcases": testcases,
		},
	))
}

// SubmitCodeHandler accepts program code blocks to run evaluation processes.
func SubmitCodeHandler(c *gin.Context) {

	var req SubmitCodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	submission, err := SubmitCodeService(
		req,
		user.ID,
		true,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusAccepted, utils.NewResponse(
		"submission queued successfully",
		gin.H{
			"submission": submission,
		},
	))
}

// SubmitSampleCodeHandler checks input blocks against primary sample testcases.
func SubmitSampleCodeHandler(c *gin.Context) {

	var req SubmitCodeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid request payload",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	submission, err := SubmitCodeService(
		req,
		user.ID,
		false,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusAccepted, utils.NewResponse(
		"sample submission queued successfully",
		gin.H{
			"submission": submission,
		},
	))
}

// GetSubmissionStatusHandler responds with processing details on code submission status.
func GetSubmissionStatusHandler(c *gin.Context) {

	submissionID, err := uuid.Parse(c.Param("submissionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid submission id format",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if err = VerifySubmissionOwnership(submissionID,
		user.ID); err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	submission, err := GetSubmissionStatusService(
		submissionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"submission status fetched successfully",
		gin.H{
			"submission": submission,
		},
	))
}

// GetSubmissionResultHandler details score values once processing settles.
func GetSubmissionResultHandler(c *gin.Context) {

	submissionID, err := uuid.Parse(c.Param("submissionId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid submission id format",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	if err = VerifySubmissionOwnership(submissionID,
		user.ID); err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	results, err := GetSubmissionResultService(
		submissionID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"submission results fetched successfully",
		gin.H{
			"results": results,
		},
	))
}
