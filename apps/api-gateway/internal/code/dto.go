// Package code manages structural data formatting for coding tasks.
package code

import (
	"time"

	"github.com/google/uuid"
)

// CreateCodingPracticeRequest holds creation details for a practice set.
type CreateCodingPracticeRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
	PlaylistID  uuid.UUID `json:"playlist_id" binding:"required"`
}

// CreateCodingQuestionRequest holds creation details for a coding challenge.
type CreateCodingQuestionRequest struct {
	ContestID     uuid.UUID `json:"contest_id" binding:"required"`
	Title         string    `json:"title" binding:"required"`
	Difficulty    string    `json:"difficulty" binding:"required"`
	Statement     string    `json:"statement" binding:"required"`
	Constraints   *string   `json:"constraints"`
	InputFormat   *string   `json:"input_format"`
	OutputFormat  *string   `json:"output_format"`
	TimeLimitMS   int32     `json:"time_limit_ms" binding:"required"`
	MemoryLimitMB int32     `json:"memory_limit_mb" binding:"required"`
}

// CreateTestCaseRequest holds creation details for a question test case.
type CreateTestCaseRequest struct {
	QuestionID     uuid.UUID `json:"question_id" binding:"required"`
	Input          string    `json:"input" binding:"required"`
	ExpectedOutput string    `json:"expected_output" binding:"required"`
	IsHidden       bool      `json:"is_hidden"`
}

// SubmitCodeRequest holds data submitted by a user for validation.
type SubmitCodeRequest struct {
	QuestionID uuid.UUID `json:"question_id" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	Language   string    `json:"language" binding:"required"`
}

// CodingQuestionMetadata holds information describing a question.
type CodingQuestionMetadata struct {
	ID         uuid.UUID `json:"question_id"`
	Title      string    `json:"title"`
	Difficulty string    `json:"difficulty"`
}

// CodingPracticeResponse holds data sent back for a practice set.
type CodingPracticeResponse struct {
	ID          uuid.UUID                `json:"contest_id"`
	Title       string                   `json:"title"`
	Description *string                  `json:"description"`
	PlaylistID  uuid.UUID                `json:"playlist_id"`
	CreatedAt   time.Time                `json:"created_at"`
	Questions   []CodingQuestionMetadata `json:"questions"`
}

// SubmissionStatusResponse holds general updates on a code submission.
type SubmissionStatusResponse struct {
	ID          uuid.UUID  `json:"submission_id"`
	Status      string     `json:"status"`
	PassedCases int32      `json:"passed_cases"`
	TotalCases  int32      `json:"total_cases"`
	StartedAt   *time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`
}

// SubmissionResultResponse holds final evaluation results for code.
type SubmissionResultResponse struct {
	SubmissionID   uuid.UUID `json:"submission_id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expected_output"`
	ActualOutput   *string   `json:"actual_output"`
	ErrorOutput    *string   `json:"error_output"`
	IsPassed       bool      `json:"is_passed"`
	Verdict        string    `json:"verdict"`
	RuntimeMS      *int32    `json:"runtime_ms"`
	MemoryKB       *int32    `json:"memory_kb"`
}
