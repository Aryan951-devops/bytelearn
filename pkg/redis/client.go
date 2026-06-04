package redis

import (
	"context"
	"log"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type Client struct {
	*goredis.Client
}

func (c *Client) Ping(
	ctx context.Context,
) error {

	return c.Client.Ping(ctx).Err()
}

func New(addr string,
	password string,
	db int,
) *Client {

	return &Client{
		Client: goredis.NewClient(
			&goredis.Options{
				Addr:     addr,
				Password: password,
				DB:       db,
			},
		),
	}
}

func MustConnectRedis(addr string,
	password string,
	db int,
) *Client {

	client := New(addr, password, db)
	ctx, cancel :=
		context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
	defer cancel()

	err := client.Ping(ctx)

	if err != nil {
		log.Fatalf(
			"redis connection failed: %v",
			err,
		)
	}

	log.Print("Ping Success!")

	return client
}
