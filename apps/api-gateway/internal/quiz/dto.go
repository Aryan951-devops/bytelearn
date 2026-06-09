package quiz

import (
	"time"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

type CreateQuizRequest struct {
	Title           string               `json:"title" binding:"required"`
	PlaylistID      uuid.UUID            `json:"playlist_id" binding:"required"`
	DurationMinutes int32                `json:"duration_minutes" binding:"required"`
	Questions       []CreateQuizQuestion `json:"questions" binding:"required"`
}

type CreateQuizQuestion struct {
	Type           models.QuestionType `json:"type"`
	Question       string              `json:"question"`
	Options        []string            `json:"options"`
	CorrectOptions []int               `json:"correct_options"`
	CorrectAnswer  string              `json:"correct_answer"`
	Marks          int32               `json:"marks"`
	NegativeMarks  int32               `json:"negative_marks"`
	Explanation    *string             `json:"explanation"`
}

type QuizQuestionMetadata struct {
	ID            uuid.UUID           `json:"question_id"`
	Type          models.QuestionType `json:"type"`
	Question      string              `json:"question"`
	Options       []string            `json:"options,omitempty"`
	Marks         int32               `json:"marks"`
	NegativeMarks int32               `json:"negative_marks"`
}

type StartQuizResponse struct {
	AttemptID       uuid.UUID              `json:"attempt_id"`
	QuizID          uuid.UUID              `json:"quiz_id"`
	DurationMinutes int32                  `json:"duration_minutes"`
	StartedAt       time.Time              `json:"started_at"`
	Questions       []QuizQuestionMetadata `json:"questions"`
}

type SubmitQuizRequest struct {
	Answers []models.UserSubmittedAnswer `json:"answers"`
}

type SubmitQuizResponse struct {
	ID         uuid.UUID                `json:"attempt_id"`
	Score      int32                    `json:"score"`
	TotalMarks int32                    `json:"total_marks"`
	Status     models.QuizAttemptStatus `json:"status"`
}

type QuizAttemptResponse struct {
	ID               uuid.UUID            `json:"attempt_id"`
	QuizID           uuid.UUID            `json:"quiz_id"`
	UserID           uuid.UUID            `json:"user_id"`
	Score            int32                `json:"score"`
	TotalMarks       int32                `json:"total_marks"`
	SubmittedAnswers []QuizAnswerResponse `json:"submitted_answers"`
	StartedAt        time.Time            `json:"started_at"`
	SubmittedAt      time.Time            `json:"submitted_at"`
}

type QuizAnswerResponse struct {
	QuestionID     uuid.UUID           `json:"question_id"`
	Type           models.QuestionType `json:"type"`
	Question       string              `json:"question"`
	Options        []string            `json:"options,omitempty"`
	SelectedOption []int               `json:"selected_options,omitempty"` // For MCQ/Multiple choice
	TextAnswer     string              `json:"text_answer,omitempty"`      // For OneWord/TrueFalse
	CorrectOptions []int               `json:"correct_options,omitempty"`
	CorrectAnswer  string              `json:"correct_answer,omitempty"`
	Explanation    *string             `json:"explanation,omitempty"`
	Marks          int32               `json:"marks"`
	NegativeMarks  int32               `json:"negative_marks"`
}
