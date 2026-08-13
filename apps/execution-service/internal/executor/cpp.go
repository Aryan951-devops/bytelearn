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

func RunCppContainer(submissionData *utils.SubmissionData,
	testcases []models.TestCase,
) ([]models.SubmissionResults, error) {

	workspace := filepath.Join(
		os.TempDir(),
		"judge-workspaces",
		"submission-"+uuid.NewString(),
	)

	if err := os.MkdirAll(workspace, 0755); err != nil {
		return nil, err
	}

	defer os.RemoveAll(workspace)

	sourceFile := filepath.Join(
		workspace,
		"main.cpp",
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

		"judge-cpp",

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

	compileCtx, compileCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer compileCancel()

	compileCmd := exec.CommandContext(
		compileCtx,
		"docker", "exec",
		containerName,
		"g++", "-O2", "/workspace/main.cpp", "-o", "/workspace/solution",
	)

	var compileStderr bytes.Buffer
	compileCmd.Stderr = &compileStderr

	if err := compileCmd.Run(); err != nil {
		compileErrStr := strings.TrimSpace(compileStderr.String())
		log.Printf("Compilation Error for submission %s: %s", submissionData.ID, compileErrStr)

		// Return Compilation Error verdict for all test cases
		results := make([]models.SubmissionResults, 0, len(testcases))
		for _, tc := range testcases {
			results = append(results, models.SubmissionResults{
				SubmissionID: submissionData.ID,
				TestCaseID:   tc.ID,
				ErrorOutput:  &compileErrStr,
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

			"/workspace/solution",
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
