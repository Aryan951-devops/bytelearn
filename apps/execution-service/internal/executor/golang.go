package executor

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/utils"
	"github.com/Aryan951-devops/bytelearn/pkg/constants"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/google/uuid"
)

func RunGoContainer(submissionData *utils.SubmissionData,
	testcases []models.TestCase,
) ([]models.SubmissionResults, error) {

	workspace := filepath.Join(
		os.TempDir(),
		"judge-workspaces",
		"submission-"+uuid.NewString(),
	)

	if err := os.MkdirAll(workspace, 0777); err != nil {
		return nil, err
	}

	defer os.RemoveAll(workspace)

	sourceFile := filepath.Join(
		workspace,
		"main.go",
	)

	if err := os.WriteFile(
		sourceFile,
		[]byte(submissionData.Code),
		0644,
	); err != nil {
		return nil, err
	}

	containerName := "judge-" + uuid.NewString()
	log.Println("Creating docker container: ", containerName)

	createCmd := exec.Command(
		"docker",
		"run",
		"-d",

		"--rm",

		"--network",
		"none",

		"--memory",
		fmt.Sprintf("%dm", submissionData.MemoryLimitMB),

		"--cpus",
		"1",

		"--pids-limit",
		"64",

		"--name",
		containerName,

		"-v",
		fmt.Sprintf("%s:/workspace", workspace),

		"judge-go",

		"tail",
		"-f",
		"/dev/null",
	)

	createOutput, err := createCmd.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to start container: %s",
			string(createOutput),
		)
	}

	log.Println("Container Created Successfully: ", containerName)

	defer func() {
		_ = exec.Command(
			"docker",
			"stop",
			containerName,
		).Run()
	}()

	buildCmd := exec.Command(
		"docker",
		"exec",
		containerName,
		"go",
		"build",
		"-o",
		"/workspace/main",
		"/workspace/main.go",
	)

	var buildStdout bytes.Buffer
	var buildStderr bytes.Buffer

	buildCmd.Stdout = &buildStdout
	buildCmd.Stderr = &buildStderr

	if buildErr := buildCmd.Run(); buildErr != nil {
		errorOutput := strings.TrimSpace(buildStderr.String())
		log.Println("Compilation failed:", errorOutput)

		results := make([]models.SubmissionResults, 0, len(testcases))
		for _, tc := range testcases {
			results = append(results, models.SubmissionResults{
				SubmissionID: submissionData.ID,
				TestCaseID:   tc.ID,
				ErrorOutput:  &errorOutput,
				IsPassed:     false,
				Verdict:      constants.VerdictCompilationError,
			})
		}
		return results, nil
	}

	results := make(
		[]models.SubmissionResults,
		0,
		len(testcases),
	)

	for _, tc := range testcases {

		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(submissionData.TimeLimitMS)*
				time.Millisecond,
		)

		start := time.Now()

		cmd := exec.CommandContext(
			ctx,

			"docker",
			"exec",
			"-i",

			containerName,

			"/workspace/main",
		)

		cmd.Stdin = strings.NewReader(tc.Input)

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		runtimeMS := int32(
			time.Since(start).Milliseconds(),
		)

		actualOutput :=
			strings.TrimSpace(
				stdout.String(),
			)

		errorOutput :=
			strings.TrimSpace(
				stderr.String(),
			)

		expectedOutput :=
			strings.TrimSpace(
				tc.ExpectedOutput,
			)

		result := models.SubmissionResults{
			SubmissionID: submissionData.ID,
			TestCaseID:   tc.ID,
			ActualOutput: &actualOutput,
			ErrorOutput:  &errorOutput,
			RuntimeMS:    &runtimeMS,
		}

		switch {

		case ctx.Err() == context.DeadlineExceeded:
			result.IsPassed = false
			result.Verdict = constants.VerdictTLE

		case err != nil:
			result.IsPassed = false
			result.Verdict = constants.VerdictRuntimeError

		case actualOutput != expectedOutput:
			result.IsPassed = false
			result.Verdict = constants.VerdictWrongAnswer

		default:
			result.IsPassed = true
			result.Verdict = constants.VerdictAccepted
		}

		results = append(results, result)

		defer cancel()
	}

	return results, nil
}
