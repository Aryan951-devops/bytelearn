package quiz

import (
	"net/http"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/playlist"
	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateQuizHandler handles the API route to save a new quiz.
func CreateQuizHandler(c *gin.Context) {

	var req CreateQuizRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	err := playlist.VerifyPlaylistOwnership(
		req.PlaylistID,
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	quiz, err := CreateQuizService(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, utils.NewResponse(
		"quiz created successfully",
		gin.H{
			"quiz": quiz,
		},
	))
}

// StartQuizHandler handles the API route to begin a quiz.
func StartQuizHandler(c *gin.Context) {

	quizID, err := uuid.Parse(c.Param("quizId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid quiz id",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	response, err := StartQuizService(quizID, user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"quiz started successfully",
		response,
	))
}

// GetAllQuizesOfPlaylistHandler handles the API route to list quizzes.
func GetAllQuizesOfPlaylistHandler(c *gin.Context) {

	playlistID, err := uuid.Parse(c.Param("playlistId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist id",
			nil,
		))
		return
	}

	quizzes, err := GetAllQuizesOfPlaylistService(playlistID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"quizzes fetched successfully",
		quizzes,
	))
}

// SubmitQuizHandler handles the API route to score a quiz.
func SubmitQuizHandler(c *gin.Context) {

	attemptID, err := uuid.Parse(c.Param("attemptId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid attempt id",
			nil,
		))
		return
	}

	var req SubmitQuizRequest

	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	err = VerifyAttemptOwnership(
		attemptID,
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
		return
	}

	result, err := SubmitQuizService(attemptID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"quiz submitted successfully",
		result,
	))
}

// GetAttemptResultHandler handles the API route to get quiz result.
func GetAttemptResultHandler(c *gin.Context) {

	attemptID, err := uuid.Parse(c.Param("attemptId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid attempt id",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	err = VerifyAttemptOwnership(
		attemptID,
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"unauthorized",
			nil,
		))
	}

	attemptResult, err := GetAttemptResultService(
		attemptID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"attempt result fetched successfully",
		attemptResult,
	))

}

// GetAllAttemptsOfQuizHandler handles the API route to see previous tries.
func GetAllAttemptsOfQuizHandler(c *gin.Context) {

	quizID, err := uuid.Parse(c.Param("quizId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewResponse(
			"invalid playlist id",
			nil,
		))
		return
	}

	ctxUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, utils.NewResponse(
			"Session user context not found",
			nil,
		))
		return
	}

	user := ctxUser.(*models.User)

	attempts, err := GetAllAttemptsOfQuizService(
		quizID,
		user.ID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.NewResponse(
			err.Error(),
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, utils.NewResponse(
		"attempts fetched successfully",
		attempts,
	))
}
