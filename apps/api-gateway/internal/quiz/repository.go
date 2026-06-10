package quiz

import (
	"context"
	"fmt"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreateQuiz saves a quiz record to the database.
func CreateQuiz(
	quiz *models.Quiz,
	questions []CreateQuizQuestion,
) (*models.Quiz, error) {

	tx, err := database.DB.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(context.Background())
	}()

	query := `
	INSERT INTO quizzes
	(title, playlist_id, duration_minutes,
	created_at, updated_at)
	VALUES ($1,$2,$3,NOW(),NOW())
	RETURNING quiz_id,title,
	playlist_id,duration_minutes,
	created_at,updated_at
	`

	var newQuiz models.Quiz

	err = tx.QueryRow(
		context.Background(),
		query,
		quiz.Title,
		quiz.PlaylistID,
		quiz.DurationMinutes,
	).Scan(
		&newQuiz.ID,
		&newQuiz.Title,
		&newQuiz.PlaylistID,
		&newQuiz.DurationMinutes,
		&newQuiz.CreatedAt,
		&newQuiz.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	for _, q := range questions {

		questionQuery := `
			INSERT INTO quiz_questions
			(
				quiz_id, type, question, options,
				correct_options, correct_answer,
				marks, negative_marks, explanation
			)
			VALUES
			($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`

		_, err = tx.Exec(
			context.Background(),
			questionQuery,
			newQuiz.ID,
			q.Type,
			q.Question,
			q.Options,
			q.CorrectOptions,
			q.CorrectAnswer,
			q.Marks,
			q.NegativeMarks,
			q.Explanation,
		)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit(context.Background())

	return &newQuiz, err
}

// CreateQuizAttempt records a user beginning a quiz.
func CreateQuizAttempt(
	quizID uuid.UUID,
	userID uuid.UUID,
) (*models.QuizAttempt, error) {

	query := `
	INSERT INTO quiz_attempts
	(quiz_id,user_id,status,started_at)
	VALUES ($1,$2,$3,NOW())
	RETURNING
		attempt_id, quiz_id, user_id,
		score, total_marks, status,
		started_at
	`

	var attempt models.QuizAttempt

	err := database.DB.QueryRow(
		context.Background(),
		query,
		quizID,
		userID,
		models.AttemptInProgress,
	).Scan(
		&attempt.ID,
		&attempt.QuizID,
		&attempt.UserID,
		&attempt.Score,
		&attempt.TotalMarks,
		&attempt.Status,
		&attempt.StartedAt,
	)

	return &attempt, err
}

// GetQuizByID finds a quiz using its ID.
func GetQuizByID(quizID uuid.UUID) (*models.Quiz, error) {

	query := `
		SELECT quiz_id, title, playlist_id,
		duration_minutes, created_at, updated_at
		FROM quizzes 
		WHERE quiz_id = $1
	`

	var quiz models.Quiz

	err := database.DB.QueryRow(
		context.Background(),
		query,
		quizID,
	).Scan(
		&quiz.ID,
		&quiz.Title,
		&quiz.PlaylistID,
		&quiz.DurationMinutes,
		&quiz.CreatedAt,
		&quiz.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &quiz, nil
}

// GetQuizDetailedQuestions fetches questions with details.
func GetQuizDetailedQuestions(quizID uuid.UUID,
) ([]models.QuizQuestion, error) {

	query := `
		SELECT question_id, type, question,
		options, correct_options, correct_answer,
		marks, negative_marks
		FROM quiz_questions 
		WHERE quiz_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		quizID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quiz_questions := []models.QuizQuestion{}

	for rows.Next() {
		var ques models.QuizQuestion

		err := rows.Scan(
			&ques.ID,
			&ques.Type,
			&ques.Question,
			&ques.Options,
			&ques.CorrectOptions,
			&ques.CorrectAnswer,
			&ques.Marks,
			&ques.NegativeMarks,
		)

		if err != nil {
			return nil, err
		}

		quiz_questions = append(quiz_questions, ques)
	}

	return quiz_questions, nil
}

// GetQuizQuestions fetches a list of questions for a quiz.
func GetQuizQuestions(quizID uuid.UUID) ([]QuizQuestionMetadata, error) {

	query := `
		SELECT question_id, type, question,
		options, marks, negative_marks
		FROM quiz_questions 
		WHERE quiz_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		quizID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quizQuestions := []QuizQuestionMetadata{}

	for rows.Next() {
		var ques QuizQuestionMetadata

		err := rows.Scan(
			&ques.ID,
			&ques.Type,
			&ques.Question,
			&ques.Options,
			&ques.Marks,
			&ques.NegativeMarks,
		)

		if err != nil {
			return nil, err
		}

		quizQuestions = append(quizQuestions, ques)
	}

	return quizQuestions, nil
}

// GetAttemptByID gets a quiz attempt file by its ID.
func GetAttemptByID(attemptID uuid.UUID,
) (*models.QuizAttempt, error) {

	query := `
	SELECT
		attempt_id, quiz_id, user_id,
		score, total_marks, status, 
		submitted_answers, started_at, submitted_at
	FROM quiz_attempts
	WHERE attempt_id = $1
	`

	var attempt models.QuizAttempt

	err := database.DB.QueryRow(
		context.Background(),
		query,
		attemptID,
	).Scan(
		&attempt.ID,
		&attempt.QuizID,
		&attempt.UserID,
		&attempt.Score,
		&attempt.TotalMarks,
		&attempt.Status,
		&attempt.SubmittedAnswers,
		&attempt.StartedAt,
		&attempt.SubmittedAt,
	)

	if err != nil {
		return nil, err
	}

	return &attempt, nil
}

// GetAttemptResultByID returns results for an attempt ID.
func GetAttemptResultByID(attemptID uuid.UUID) (*QuizAttemptResponse, error) {

	query := `
	SELECT
		attempt_id, quiz_id, user_id,
		score, total_marks, submitted_answers,
		started_at, submitted_at
	FROM quiz_attempts
	WHERE attempt_id = $1
	`

	var attempt models.QuizAttempt

	err := database.DB.QueryRow(
		context.Background(),
		query,
		attemptID,
	).Scan(
		&attempt.ID,
		&attempt.QuizID,
		&attempt.UserID,
		&attempt.Score,
		&attempt.TotalMarks,
		&attempt.SubmittedAnswers,
		&attempt.StartedAt,
		&attempt.SubmittedAt,
	)

	if err != nil {
		return nil, err
	}

	questions, err := GetQuizDetailedQuestions(attempt.QuizID)
	if err != nil {
		return nil, err
	}

	answerMap := make(map[uuid.UUID]models.UserSubmittedAnswer)

	for _, ans := range attempt.SubmittedAnswers {
		answerMap[ans.QuestionID] = ans
	}

	response := &QuizAttemptResponse{
		ID:               attempt.ID,
		QuizID:           attempt.QuizID,
		UserID:           attempt.UserID,
		Score:            attempt.Score,
		TotalMarks:       attempt.TotalMarks,
		StartedAt:        attempt.StartedAt,
		SubmittedAt:      attempt.SubmittedAt,
		SubmittedAnswers: []QuizAnswerResponse{},
	}

	for _, q := range questions {

		ans := answerMap[q.ID]

		response.SubmittedAnswers = append(
			response.SubmittedAnswers,
			QuizAnswerResponse{
				QuestionID:     q.ID,
				Type:           q.Type,
				Question:       q.Question,
				Options:        q.Options,
				SelectedOption: ans.SelectedOption,
				TextAnswer:     ans.TextAnswer,
				CorrectOptions: q.CorrectOptions,
				CorrectAnswer:  q.CorrectAnswer,
				Explanation:    q.Explanation,
				Marks:          q.Marks,
				NegativeMarks:  q.NegativeMarks,
			},
		)
	}

	return response, nil
}

// SubmitAttempt submits the attempt.
func SubmitAttempt(attempt *models.QuizAttempt,
) (*SubmitQuizResponse, error) {

	query := `
		UPDATE quiz_attempts 
		SET submitted_answers = $1,
		score = $2,
		total_marks = $3,
		submitted_at = NOW(),
		status = $4
		WHERE attempt_id = $5
		RETURNING 
		attempt_id, status, score,
		total_marks
	`

	var response SubmitQuizResponse

	err := database.DB.QueryRow(
		context.Background(),
		query,
		attempt.SubmittedAnswers,
		attempt.Score,
		attempt.TotalMarks,
		attempt.Status,
		attempt.ID,
	).Scan(
		&response.ID,
		&response.Status,
		&response.Score,
		&response.TotalMarks,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetAllQuizesOfPlaylist reads all quizzes in a playlist folder.
func GetAllQuizesOfPlaylist(playlistID uuid.UUID,
) ([]models.Quiz, error) {

	query := `
		SELECT quiz_id, title, playlist_id, 
		duration_minutes, created_at, updated_at
		FROM quizzes
		WHERE playlist_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		playlistID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	quizzes := []models.Quiz{}

	for rows.Next() {

		var quiz models.Quiz

		err = rows.Scan(
			&quiz.ID,
			&quiz.Title,
			&quiz.PlaylistID,
			&quiz.DurationMinutes,
			&quiz.CreatedAt,
			&quiz.UpdatedAt,
		)

		fmt.Println("err: ", err)
		if err != nil {
			return nil, err
		}

		quizzes = append(quizzes, quiz)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}

// GetAttemptsOfQuiz finds user's attempts on a quiz.
func GetAttemptsOfQuiz(
	quizID uuid.UUID,
	userID uuid.UUID,
) ([]models.QuizAttempt, error) {

	query := `
		SELECT 
			attempt_id, score, total_marks, 
			started_at, submitted_at
		FROM quiz_attempts
		WHERE quiz_id = $1 AND
		user_id = $2 AND
		status = 'submitted'
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		quizID,
		userID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	attempts := []models.QuizAttempt{}

	for rows.Next() {

		var attempt models.QuizAttempt

		err := rows.Scan(
			&attempt.ID,
			&attempt.Score,
			&attempt.TotalMarks,
			&attempt.StartedAt,
			&attempt.SubmittedAt,
		)

		if err != nil {
			return nil, err
		}

		attempts = append(attempts, attempt)
	}

	return attempts, nil
}

// VerifyAttemptOwnership checks if a user owns a quiz attempt record.
func VerifyAttemptOwnership(
	attemptID uuid.UUID,
	userID uuid.UUID,
) error {

	query := `
	SELECT attempt_id
	FROM quiz_attempts
	WHERE attempt_id=$1 AND user_id=$2
	`

	var id uuid.UUID

	return database.DB.QueryRow(
		context.Background(),
		query,
		attemptID,
		userID,
	).Scan(&id)
}

// GetQuizDurationByID looks up the time limit for a quiz.
func GetQuizDurationByID(
	quizID uuid.UUID,
) (int32, error) {

	query := `
		SELECT duration_minutes 
		FROM quizzes
		WHERE quiz_id = $1
	`

	var timeDuration int32
	err := database.DB.QueryRow(
		context.Background(),
		query,
		quizID,
	).Scan(&timeDuration)

	if err != nil {
		return 0, err
	}

	return timeDuration, nil
}
