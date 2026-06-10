package code

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/config"
	"github.com/Aryan951-devops/bytelearn/pkg/constants"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreateCodingPracticeService prepares data for storing practice elements.
func CreateCodingPracticeService(
	req CreateCodingPracticeRequest,
) (*models.CodingPractice, error) {

	practice := models.CodingPractice{
		Title:       req.Title,
		Description: req.Description,
		PlaylistID:  req.PlaylistID,
	}

	return CreateCodingPractice(&practice)
}

// GetCodingPracticeService retrives coding practice.
func GetCodingPracticeService(contestID uuid.UUID,
) (*CodingPracticeResponse, error) {

	return GetCodingPracticeByID(contestID)
}

// GetCodingPracticesOfPlaylistService retrives all practices of playlist.
func GetCodingPracticesOfPlaylistService(playlistID uuid.UUID,
) (*[]models.CodingPractice, error) {

	return GetCodingPracticesOfPlaylist(playlistID)
}

// GetCodingQuestionService finds a problem configuration using its ID.
func GetCodingQuestionService(questionID uuid.UUID,
) (*models.CodingQuestion, error) {

	return GetCodingQuestionByID(questionID)
}

// CreateCodingQuestionService creates a coding question.
func CreateCodingQuestionService(
	req CreateCodingQuestionRequest,
) (*models.CodingQuestion, error) {

	question := models.CodingQuestion{
		ContestID:     req.ContestID,
		Title:         req.Title,
		Difficulty:    req.Difficulty,
		Statement:     req.Statement,
		Constraints:   req.Constraints,
		InputFormat:   req.InputFormat,
		OutputFormat:  req.OutputFormat,
		TimeLimitMS:   req.TimeLimitMS,
		MemoryLimitMB: req.MemoryLimitMB,
	}

	return CreateCodingQuestion(&question)
}

// CreateCodingTestCaseService sets up individual testing logic rules.
func CreateCodingTestCaseService(
	req CreateTestCaseRequest,
) (*models.TestCase, error) {

	testcase := models.TestCase{
		QuestionID:     req.QuestionID,
		Input:          req.Input,
		ExpectedOutput: req.ExpectedOutput,
		IsHidden:       req.IsHidden,
	}

	return CreateCodingTestCase(&testcase)
}

// GetSampleTestCasesService manages lookup requests for publicly visible examples.
func GetSampleTestCasesService(questionID uuid.UUID,
) (*[]models.TestCase, error) {

	return GetSampleTestCases(questionID)
}

// SubmitCodeService processes incoming text solutions to be checked.
func SubmitCodeService(req SubmitCodeRequest,
	userID uuid.UUID,
	isHidden bool,
) (*models.Submission, error) {

	submission := models.Submission{
		QuestionID: req.QuestionID,
		Code:       req.Code,
		Language:   req.Language,
		UserID:     userID,
	}

	newSubmission, err := CreateSubmission(&submission)

	if err != nil {
		return nil, err
	}

	job := models.SubmissionJob{
		SubmissionID: newSubmission.ID,
		IsHidden:     isHidden,
	}

	err = config.AppConfig.RedisClient.Publish(
		context.Background(),
		constants.SubmissionQueue,
		job,
	)

	if err != nil {
		return nil, err
	}

	return newSubmission, nil
}

// GetSubmissionStatusService monitors background execution lines.
func GetSubmissionStatusService(submissionID uuid.UUID,
) (*SubmissionStatusResponse, error) {

	return FetchSubmissionStatus(submissionID)
}

// GetSubmissionResultService gathers score parameters once runs complete.
func GetSubmissionResultService(submissionID uuid.UUID,
) (*[]SubmissionResultResponse, error) {

	return FetchSubmissionResult(submissionID)
}
