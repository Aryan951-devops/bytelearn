package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/config"
	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/database"
	"github.com/Aryan951-devops/bytelearn/apps/execution_service/internal/worker"
	"github.com/Aryan951-devops/bytelearn/pkg/constants"
	"github.com/Aryan951-devops/bytelearn/pkg/models"
	"github.com/Aryan951-devops/bytelearn/pkg/redis"
)

func main() {

	config.LoadConfig()
	database.ConnectDB(config.AppConfig.DatabaseURL)

	redisClient := redis.MustConnectRedis(
		config.AppConfig.RedisAddr,
		"",
		0,
	)

	jobs := make(
		chan *models.SubmissionJob,
		100,
	)

	ctx := context.Background()

	go func() {

		for {

			job, err := redisClient.Consume(
				ctx,
				constants.SubmissionQueue,
			)

			if err != nil {
				log.Println(err)
				time.Sleep(1 * time.Second)
				continue
			}

			jobs <- job
		}
	}()

	workerCount := runtime.NumCPU()

	for i := 0; i < workerCount; i++ {

		go worker.Start(
			ctx,
			jobs,
		)
	}

	select {}
}
