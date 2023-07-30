package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/salesforceanton/go-worker-pool/activity_generator"
	"github.com/salesforceanton/go-worker-pool/workerpool"
)

const (
	WORKERS_COUNT = 10
	USERS_COUNT   = 100
)

func main() {
	rand.Seed(time.Now().Unix())

	startTime := time.Now()

	results := make(chan workerpool.SaveUserLogsResult, USERS_COUNT)
	workerPool := workerpool.New(WORKERS_COUNT, USERS_COUNT, results)

	// Run worker goroutines which wait for a jobs
	workerPool.Init()

	// Push tasks to queue and read results in parallel
	generateJobs(workerPool)
	proccessResults(results)

	fmt.Printf("DONE! Time Elapsed: %.2f seconds\n", time.Since(startTime).Seconds())
}

func proccessResults(results chan workerpool.SaveUserLogsResult) {
	for i := 0; i < USERS_COUNT; i++ {
		result := <-results
		fmt.Println(result.Info())
	}
}

func generateJobs(pool *workerpool.Pool) {
	for i := 0; i < USERS_COUNT; i++ {
		pool.Push(
			workerpool.SaveUserLogsJob{
				User: activity_generator.GenerateUser(i),
			},
		)
	}
	pool.Stop()
}
