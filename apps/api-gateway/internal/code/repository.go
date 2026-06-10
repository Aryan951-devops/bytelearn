package code

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

// CreateCodingPractice saves a coding practice set.
func CreateCodingPractice(practice *models.CodingPractice,
) (*models.CodingPractice, error) {

	query := `
		INSERT into coding_practice
		(title, description, playlist_id)
		VALUES
		($1,$2,$3)
		RETURNING
		contest_id, title, description,
		playlist_id, created_at, updated_at
	`

	var newPractice models.CodingPractice

	err := database.DB.QueryRow(
		context.Background(),
		query,
		practice.Title,
		practice.Description,
		practice.PlaylistID,
	).Scan(
		&newPractice.ID,
		&newPractice.Title,
		&newPractice.Description,
		&newPractice.PlaylistID,
		&newPractice.CreatedAt,
		&newPractice.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newPractice, nil
}

// GetCodingPracticeByID retrieves a practice set by its ID.
func GetCodingPracticeByID(contestID uuid.UUID,
) (*CodingPracticeResponse, error) {

	query := `
		SELECT contest_id, title, description,
		playlist_id, created_at
		FROM coding_practice
		WHERE contest_id = $1
	`

	practice := CodingPracticeResponse{}
	practice.Questions = []CodingQuestionMetadata{}

	err := database.DB.QueryRow(
		context.Background(),
		query,
		contestID,
	).Scan(
		&practice.ID,
		&practice.Title,
		&practice.Description,
		&practice.PlaylistID,
		&practice.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	questionQuery := `
		SELECT question_id, title, difficulty
		FROM coding_questions
		WHERE contest_id = $1
	`

	rows, err := database.DB.Query(
		context.Background(),
		questionQuery,
		contestID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var question CodingQuestionMetadata

		err := rows.Scan(
			&question.ID,
			&question.Title,
			&question.Difficulty,
		)

		if err != nil {
			return nil, err
		}

		practice.Questions = append(practice.Questions, question)
	}

	return &practice, nil
}

// GetCodingPracticesOfPlaylist fetches lists matching a folder identifier.
func GetCodingPracticesOfPlaylist(playlistID uuid.UUID,
) (*[]models.CodingPractice, error) {

	query := `
		SELECT contest_id, title, description,
		playlist_id, created_at, updated_at
		FROM coding_practice
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

	cps := []models.CodingPractice{}

	for rows.Next() {
		var cp models.CodingPractice

		err := rows.Scan(
			&cp.ID,
			&cp.Title,
			&cp.Description,
			&cp.PlaylistID,
			&cp.CreatedAt,
			&cp.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		cps = append(cps, cp)
	}

	return &cps, nil
}

// GetCodingQuestionByID retrieves a coding problem configuration.
func GetCodingQuestionByID(questionID uuid.UUID,
) (*models.CodingQuestion, error) {
	query := `
		SELECT
			question_id, contest_id, title, difficulty,
			statement, constraints, input_format,
			output_format, time_limit_ms, memory_limit_mb,
			created_at, updated_at
		FROM coding_questions
		WHERE question_id = $1
	`

	var question models.CodingQuestion

	err := database.DB.QueryRow(
		context.Background(),
		query,
		questionID,
	).Scan(
		&question.ID,
		&question.ContestID,
		&question.Title,
		&question.Difficulty,
		&question.Statement,
		&question.Constraints,
		&question.InputFormat,
		&question.OutputFormat,
		&question.TimeLimitMS,
		&question.MemoryLimitMB,
		&question.CreatedAt,
		&question.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &question, nil
}

// CreateCodingQuestion saves a new coding question.
func CreateCodingQuestion(question *models.CodingQuestion,
) (*models.CodingQuestion, error) {

	query := `
		INSERT INTO coding_questions
		(contest_id, title, difficulty, statement,
		constraints, input_format, output_format,
		time_limit_ms, memory_limit_mb)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		RETURNING
		question_id, contest_id, title, difficulty,
		statement, constraints, input_format,
		output_format, time_limit_ms, memory_limit_mb,
		created_at, updated_at
	`

	var newQuestion models.CodingQuestion

	err := database.DB.QueryRow(
		context.Background(),
		query,
		question.ContestID,
		question.Title,
		question.Difficulty,
		question.Statement,
		question.Constraints,
		question.InputFormat,
		question.OutputFormat,
		question.TimeLimitMS,
		question.MemoryLimitMB,
	).Scan(
		&newQuestion.ID,
		&newQuestion.ContestID,
		&newQuestion.Title,
		&newQuestion.Difficulty,
		&newQuestion.Statement,
		&newQuestion.Constraints,
		&newQuestion.InputFormat,
		&newQuestion.OutputFormat,
		&newQuestion.TimeLimitMS,
		&newQuestion.MemoryLimitMB,
		&newQuestion.CreatedAt,
		&newQuestion.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newQuestion, nil
}

// CreateCodingTestCase saves validation rules for testing code solutions.
func CreateCodingTestCase(testcase *models.TestCase,
) (*models.TestCase, error) {

	query := `
		INSERT into testcases
		(question_id, input, expected_output,
		is_hidden)
		VALUES
		($1,$2,$3,$4)
		RETURNING
		testcase_id, question_id, input,
		expected_output, is_hidden, created_at
	`

	var newTestcase models.TestCase

	err := database.DB.QueryRow(
		context.Background(),
		query,
		testcase.QuestionID,
		testcase.Input,
		testcase.ExpectedOutput,
		testcase.IsHidden,
	).Scan(
		&newTestcase.ID,
		&newTestcase.QuestionID,
		&newTestcase.Input,
		&newTestcase.ExpectedOutput,
		&newTestcase.IsHidden,
		&newTestcase.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newTestcase, nil
}

// GetSampleTestCases pulls visible validation elements matching a question ID.
func GetSampleTestCases(questionID uuid.UUID,
) (*[]models.TestCase, error) {

	query := `
		SELECT 
		testcase_id, question_id, input,
		expected_output, is_hidden, created_at
		FROM testcases
		WHERE question_id = $1 AND is_hidden = false
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
		questionID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	testcases := []models.TestCase{}

	for rows.Next() {
		var tc models.TestCase

		err := rows.Scan(
			&tc.ID,
			&tc.QuestionID,
			&tc.Input,
			&tc.ExpectedOutput,
			&tc.IsHidden,
			&tc.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		testcases = append(testcases, tc)
	}

	return &testcases, nil
}

// CreateSubmission saves a code submission attempt.
func CreateSubmission(submission *models.Submission,
) (*models.Submission, error) {

	query := `
		INSERT into submissions
		(question_id, user_id, code, language)
		VALUES ($1,$2,$3,$4)
		RETURNING 
		submission_id, question_id, user_id, code,
		language, status, passed_cases, total_cases,
		started_at, finished_at, submitted_at,
		updated_at
	`

	var newSubmission models.Submission

	err := database.DB.QueryRow(
		context.Background(),
		query,
		submission.QuestionID,
		submission.UserID,
		submission.Code,
		submission.Language,
	).Scan(
		&newSubmission.ID,
		&newSubmission.QuestionID,
		&newSubmission.UserID,
		&newSubmission.Code,
		&newSubmission.Language,
		&newSubmission.Status,
		&newSubmission.PassedCases,
		&newSubmission.TotalCases,
		&newSubmission.Started_At,
		&newSubmission.Finished_At,
		&newSubmission.SubmittedAt,
		&newSubmission.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &newSubmission, nil
}

// FetchSubmissionStatus reads current verification workflow records.
func FetchSubmissionStatus(submissionID uuid.UUID,
) (*SubmissionStatusResponse, error) {

	query := `
		SELECT submission_id, status, passed_cases, 
		total_cases, started_at, finished_at 
		FROM submissions 
		WHERE submission_id = $1
	`

	var response SubmissionStatusResponse

	err := database.DB.QueryRow(
		context.Background(),
		query,
		submissionID,
	).Scan(
		&response.ID,
		&response.Status,
		&response.PassedCases,
		&response.TotalCases,
		&response.StartedAt,
		&response.FinishedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

// FetchSubmissionResult logs evaluation reports matching an ID.
func FetchSubmissionResult(submissionID uuid.UUID,
) (*[]SubmissionResultResponse, error) {

	query := `
		SELECT 
			sr.submission_id, t.input, t.expected_output, 
			sr.actual_output, sr.error_output, sr.is_passed, 
			sr.verdict, sr.runtime_ms, sr.memory_kb
		FROM submission_results sr 
			INNER JOIN testcases t
		ON sr.testcase_id = t.testcase_id
		WHERE sr.submission_id = $1

	`

	response := []SubmissionResultResponse{}

	rows, err := database.DB.Query(
		context.Background(),
		query,
		submissionID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result SubmissionResultResponse

		err := rows.Scan(
			&result.SubmissionID,
			&result.Input,
			&result.ExpectedOutput,
			&result.ActualOutput,
			&result.ErrorOutput,
			&result.IsPassed,
			&result.Verdict,
			&result.RuntimeMS,
			&result.MemoryKB,
		)

		if err != nil {
			return nil, err
		}

		response = append(response, result)
	}

	return &response, nil
}

// VerifySubmissionOwnership ensures a user owns a given submission.
func VerifySubmissionOwnership(submissionID uuid.UUID,
	userID uuid.UUID,
) error {

	query := `
	SELECT submission_id
	FROM submissions
	WHERE submission_id=$1 AND user_id=$2
	`

	var id uuid.UUID

	return database.DB.QueryRow(
		context.Background(),
		query,
		submissionID,
		userID,
	).Scan(&id)
}
