package quiz

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

var (
	// ErrQuizAlreadySubmitted is returned when a user tries to submit a quiz again.
	ErrQuizAlreadySubmitted = errors.New("quiz already submitted")

	// ErrAttemptExpired is returned when a quiz run expires.
	ErrAttemptExpired = errors.New("quiz attempt has expired, please start a new one")
)

// VerifyAndProcessAttemptState checks and updates a quiz attempt.
func VerifyAndProcessAttemptState(attempt *models.QuizAttempt,
) (*models.QuizAttempt, error) {

	switch attempt.Status {
	case models.AttemptSubmitted:
		return nil, ErrQuizAlreadySubmitted

	case models.AttemptExpired:
		return nil, ErrAttemptExpired

	default:
		timeDuration, err := GetQuizDurationByID(attempt.QuizID)
		if err != nil {
			return nil, fmt.Errorf("failed to verify quiz duration: %w", err)
		}

		elapsedMinutes := time.Since(attempt.StartedAt).Minutes()

		if elapsedMinutes > float64(timeDuration) {
			attempt.Status = models.AttemptExpired
			attempt.SubmittedAt = time.Now()

			if _, err := SubmitAttempt(attempt); err != nil {
				return nil, fmt.Errorf("failed to auto-expire attempt: %w", err)
			}

			return nil, ErrAttemptExpired
		}
	}

	return attempt, nil
}

// CreateQuizService handles creating a new quiz.
func CreateQuizService(req CreateQuizRequest,
) (*models.Quiz, error) {

	quiz := &models.Quiz{
		Title:           req.Title,
		PlaylistID:      req.PlaylistID,
		DurationMinutes: req.DurationMinutes,
	}

	return CreateQuiz(quiz, req.Questions)
}

// StartQuizService records when a user starts a quiz.
func StartQuizService(
	quizID uuid.UUID,
	userID uuid.UUID,
) (*StartQuizResponse, error) {

	attempt, err := CreateQuizAttempt(
		quizID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	quiz, err := GetQuizByID(quizID)
	if err != nil {
		return nil, err
	}

	questions, err := GetQuizQuestions(quizID)
	if err != nil {
		return nil, err
	}

	return &StartQuizResponse{
		AttemptID:       attempt.ID,
		QuizID:          quiz.ID,
		DurationMinutes: quiz.DurationMinutes,
		StartedAt:       attempt.StartedAt,
		Questions:       questions,
	}, nil
}

// SubmitQuizService saves answers and scores a quiz submission.
func SubmitQuizService(
	attemptID uuid.UUID,
	req SubmitQuizRequest,
) (*SubmitQuizResponse, error) {

	attempt, err := GetAttemptByID(attemptID)
	if err != nil {
		return nil, err
	}
	attempt, err = VerifyAndProcessAttemptState(attempt)
	if err != nil {
		return nil, err
	}

	questions, err := GetQuizDetailedQuestions(attempt.QuizID)
	if err != nil {
		return nil, err
	}

	score := int32(0)
	total := int32(0)

	for _, q := range questions {

		total += q.Marks
		for _, ans := range req.Answers {

			if ans.QuestionID != q.ID {
				continue
			}

			switch q.Type {

			case models.QuestionTypeMCQ,
				models.QuestionTypeTrueFalse,
				models.QuestionTypeMultiple:

				if len(ans.SelectedOption) > 0 && len(q.CorrectOptions) > 0 {

					flag := false
					sort.Slice(ans.SelectedOption, func(i, j int) bool {
						return i < j
					})
					sort.Slice(q.CorrectOptions, func(i, j int) bool {
						return i < j
					})
					for idx := range q.CorrectOptions {
						if ans.SelectedOption[idx] != q.CorrectOptions[idx] {
							flag = true
							break
						}
					}

					if flag {
						score -= q.NegativeMarks
					} else {
						score += q.Marks
					}

				} else {
					score -= q.NegativeMarks
				}

			case models.QuestionTypeOneWord:

				if ans.TextAnswer == q.CorrectAnswer {
					score += q.Marks
				} else {
					score -= q.NegativeMarks
				}
			}
		}
	}

	attempt.Score = score
	attempt.TotalMarks = total
	attempt.Status = models.AttemptSubmitted
	attempt.SubmittedAnswers = req.Answers
	attempt.SubmittedAt = time.Now()

	return SubmitAttempt(attempt)
}

// GetAttemptResultService gets the final result for a quiz attempt.
func GetAttemptResultService(attemptID uuid.UUID,
) (*QuizAttemptResponse, error) {

	return GetAttemptResultByID(attemptID)
}

// GetAllQuizesOfPlaylistService gets all quizzes linked to a playlist.
func GetAllQuizesOfPlaylistService(
	playlistID uuid.UUID,
) ([]models.Quiz, error) {

	return GetAllQuizesOfPlaylist(playlistID)
}

// GetAllAttemptsOfQuizService gets all history attempts of a user for a quiz.
func GetAllAttemptsOfQuizService(
	quizID uuid.UUID,
	userID uuid.UUID,
) ([]models.QuizAttempt, error) {

	return GetAttemptsOfQuiz(
		quizID,
		userID,
	)
}
