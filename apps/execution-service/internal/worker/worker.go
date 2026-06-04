package worker

import (
	"context"
	"log"

	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/executor"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
)

func Start(
	ctx context.Context,
	jobs <-chan *models.SubmissionJob,
) {

	for {

		select {

		case <-ctx.Done():
			return

		case job, ok := <-jobs:

			if !ok {
				return
			}

			err := executor.ExecuteSubmission(
				ctx,
				job,
			)

			if err != nil {
				log.Println(err)
			}
		}
	}
}
