package redis

import (
	"context"
	"encoding/json"
)

func (c *Client) Publish(
	ctx context.Context,
	queue string,
	job any,
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
) ([]byte, error) {

	result, err := c.BRPop(
		ctx,
		0,
		queue,
	).Result()

	if err != nil {
		return nil, err
	}

	return []byte(result[1]), nil
}
