package code

import (
	"context"

	"github.com/Aryan951-devops/bytelearn/apps/api-gateway/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

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

	var new_practice models.CodingPractice

	err := database.DB.QueryRow(
		context.Background(),
		query,
		practice.Title,
		practice.Description,
		practice.PlaylistID,
	).Scan(
		&new_practice.ID,
		&new_practice.Title,
		&new_practice.Description,
		&new_practice.PlaylistID,
		&new_practice.CreatedAt,
		&new_practice.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &new_practice, nil
}

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

	var new_question models.CodingQuestion

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
		&new_question.ID,
		&new_question.ContestID,
		&new_question.Title,
		&new_question.Difficulty,
		&new_question.Statement,
		&new_question.Constraints,
		&new_question.InputFormat,
		&new_question.OutputFormat,
		&new_question.TimeLimitMS,
		&new_question.MemoryLimitMB,
		&new_question.CreatedAt,
		&new_question.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &new_question, nil
}

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

	var new_testcase models.TestCase

	err := database.DB.QueryRow(
		context.Background(),
		query,
		testcase.QuestionID,
		testcase.Input,
		testcase.ExpectedOutput,
		testcase.IsHidden,
	).Scan(
		&new_testcase.ID,
		&new_testcase.QuestionID,
		&new_testcase.Input,
		&new_testcase.ExpectedOutput,
		&new_testcase.IsHidden,
		&new_testcase.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &new_testcase, nil
}

func GetSampleTestCases(question_id uuid.UUID,
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
		question_id,
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

	var new_submission models.Submission

	err := database.DB.QueryRow(
		context.Background(),
		query,
		submission.QuestionID,
		submission.UserID,
		submission.Code,
		submission.Language,
	).Scan(
		&new_submission.ID,
		&new_submission.QuestionID,
		&new_submission.UserID,
		&new_submission.Code,
		&new_submission.Language,
		&new_submission.Status,
		&new_submission.PassedCases,
		&new_submission.TotalCases,
		&new_submission.Started_At,
		&new_submission.Finished_At,
		&new_submission.SubmittedAt,
		&new_submission.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &new_submission, nil
}

func FetchSubmissionStatus(submission_id uuid.UUID,
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
		submission_id,
	).Scan(
		&response.ID,
		&response.Status,
		&response.PassedCases,
		&response.TotalCases,
		&response.Started_At,
		&response.Finished_At,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func FetchSubmissionResult(submission_id uuid.UUID,
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
		submission_id,
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

func VerifySubmissionOwnership(submission_id uuid.UUID,
	user_id uuid.UUID,
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
		submission_id,
		user_id,
	).Scan(&id)
}
