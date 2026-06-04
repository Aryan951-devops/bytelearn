package executor

import (
	"context"
	"fmt"
	"log"

	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/database"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
)

func ExecuteSubmission(
	ctx context.Context,
	job *models.SubmissionJob,
) error {

	submissionID := job.SubmissionID

	submissionData, err := database.FetchSubmissionData(ctx, submissionID)

	if err != nil {
		return err
	}

	testcases, err := database.FetchTestCases(
		ctx,
		submissionData.QuestionID,
		job.IsHidden)

	if err != nil {
		return err
	}

	err = database.UpdateSubmissionStatus(
		ctx,
		submissionID,
		"RUNNING",
	)

	if err != nil {
		return err
	}

	log.Println("Executing Code with SubmissionID: ", submissionID)

	var results []models.SubmissionResults

	switch submissionData.Language {

	case "python":

		results, err = RunPythonContainer(
			submissionData,
			testcases,
		)

	default:

		return fmt.Errorf(
			"unsupported language: %s",
			submissionData.Language,
		)
	}

	if err != nil {

		_ = database.UpdateSubmissionStatus(
			ctx,
			submissionID,
			"FAILED",
		)

		return err
	}

	log.Println("Bulk Inserting")

	err = database.BulkInsertSubmissionResults(
		ctx,
		results,
	)

	if err != nil {

		_ = database.UpdateSubmissionStatus(
			ctx,
			submissionID,
			"FAILED",
		)

		return err
	}

	var passedTestCases, totalTestCases int

	for _, tc := range results {
		if tc.IsPassed != false {
			passedTestCases += 1
		}
		totalTestCases += 1
	}

	_ = database.UpdateSubmissionTestResults(
		ctx,
		submissionID,
		passedTestCases,
		totalTestCases,
	)

	log.Println("Code Executed for submissionID: ", submissionID)

	return nil
}
