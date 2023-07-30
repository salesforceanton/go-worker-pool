package workerpool

import (
	"fmt"
	"log"

	"github.com/salesforceanton/go-worker-pool/activity_generator"
)

type SaveUserLogsResult struct {
	email        string
	filepath     string
	errorMessage string
}
type SaveUserLogsJob struct {
	User activity_generator.User
}

func (r *SaveUserLogsResult) Info() string {
	if r.errorMessage != "" {
		return fmt.Sprintf("Error with saving logs for user [%s]: %s", r.email, r.errorMessage)
	}
	return fmt.Sprintf("Logs for user [%s] have been saved successfully on: %s", r.email, r.filepath)
}

type Pool struct {
	worker       *SaveUserLogsWorker
	workersCount int

	jobs    chan SaveUserLogsJob
	results chan SaveUserLogsResult
}

func New(workersCount int, jobsCount int, results chan SaveUserLogsResult) *Pool {
	return &Pool{
		worker:       newSaveUserLogsWorker(),
		workersCount: workersCount,
		jobs:         make(chan SaveUserLogsJob, jobsCount),
		results:      results,
	}
}

func (p *Pool) Init() {
	for i := 0; i < p.workersCount; i++ {
		go p.initWorker(i)
	}
}

func (p *Pool) Push(j SaveUserLogsJob) {
	p.jobs <- j
}

func (p *Pool) Stop() {
	close(p.jobs)
}

func (p *Pool) initWorker(id int) {
	for job := range p.jobs {
		p.results <- p.worker.processJob(job)
	}

	log.Printf("[worker ID %d] finished proccesing", id)
}
