package database

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(databaseURL string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("Failed to create database pool:", err)
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = dbpool

	log.Println("Database connected successfully")
}

func FetchSubmissionData(ctx context.Context,
	submissionID uuid.UUID,
) (*utils.SubmissionData, error) {

	query := `
			SELECT s.submission_id, s.question_id, s.code, 
			s.language, cq.time_limit_ms, cq.memory_limit_mb
			FROM submissions s INNER JOIN coding_questions cq
			ON s.question_id = cq.question_id
			WHERE s.submission_id = $1
		`

	var submissionData utils.SubmissionData

	err := DB.QueryRow(
		ctx,
		query,
		submissionID,
	).Scan(
		&submissionData.ID,
		&submissionData.QuestionID,
		&submissionData.Code,
		&submissionData.Language,
		&submissionData.TimeLimitMS,
		&submissionData.MemoryLimitMB,
	)

	if err != nil {
		return nil, err
	}

	return &submissionData, nil
}

func FetchTestCases(ctx context.Context,
	questionID uuid.UUID,
	isHidden bool,
) ([]models.TestCase, error) {

	query := `
		SELECT testcase_id, question_id, input, 
		expected_output, is_hidden, created_at
		FROM testcases
		WHERE question_id = $1 AND
		is_hidden = $2
	`

	rows, err := DB.Query(
		ctx,
		query,
		questionID,
		isHidden,
	)

	if err != nil {
		return nil, errors.New(err.Error())
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
			return nil, errors.New(err.Error())
		}

		testcases = append(testcases, tc)
	}

	return testcases, nil
}

func UpdateSubmissionStatus(ctx context.Context,
	submissionID uuid.UUID,
	status string,
) error {

	query := `
		UPDATE submissions
		SET status = $1
		WHERE submission_id = $2
	`

	_, err := DB.Exec(
		ctx,
		query,
		status,
		submissionID,
	)

	return err
}

func UpdateSubmissionTestResults(ctx context.Context,
	submissionID uuid.UUID,
	passed_cases int,
	total_cases int,
) error {
	query := `
		UPDATE submissions
		SET passed_cases = $1,
			total_cases = $2,
			status = 'COMPLETED'
		WHERE submission_id = $3
	`

	_, err := DB.Exec(
		ctx,
		query,
		passed_cases,
		total_cases,
		submissionID,
	)

	return err
}

func BulkInsertSubmissionResults(
	ctx context.Context,
	results []models.SubmissionResults,
) error {

	batch := &pgx.Batch{}

	for _, result := range results {

		batch.Queue(
			`
			INSERT INTO submission_results
			(
				submission_id,
				testcase_id,
				actual_output,
				error_output,
				is_passed,
				verdict,
				runtime_ms
			)
			VALUES
			($1,$2,$3,$4,$5,$6,$7)
			`,
			result.SubmissionID,
			result.TestCaseID,
			result.ActualOutput,
			result.ErrorOutput,
			result.IsPassed,
			result.Verdict,
			result.RuntimeMS,
		)
	}

	br := DB.SendBatch(ctx, batch)

	defer br.Close()

	return nil
}
