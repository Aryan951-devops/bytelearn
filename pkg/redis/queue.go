package redis

import (
	"context"
	"encoding/json"

	"github.com/Aryan951-devops/bytelearn/pkg/models"
)

func (c *Client) Publish(
	ctx context.Context,
	queue string,
	job models.SubmissionJob,
) error {

	payload, err := json.Marshal(job)

	if err != nil {
		return err
	}

	return c.LPush(
		ctx,
		queue,
		payload,
	).Err()
}

func (c *Client) Consume(
	ctx context.Context,
	queue string,
) (*models.SubmissionJob, error) {

	result, err := c.BRPop(
		ctx,
		0,
		queue,
	).Result()

	if err != nil {
		return nil, err
	}

	var job models.SubmissionJob

	err = json.Unmarshal(
		[]byte(result[1]),
		&job,
	)

	if err != nil {
		return nil, err
	}

	return &job, nil
}
