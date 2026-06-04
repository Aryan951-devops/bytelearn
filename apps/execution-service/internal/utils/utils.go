package utils

import (
	"github.com/google/uuid"
)

type SubmissionData struct {
	ID            uuid.UUID `json:"submission_id"`
	QuestionID    uuid.UUID `json:"question_id"`
	Code          string    `json:"code"`
	Language      string    `json:"language"`
	TimeLimitMS   int32     `json:"time_limit_ms"`
	MemoryLimitMB int32     `json:"memory_limit_mb"`
}
