package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Err   error
}

func RunWorkerPool(
	ctx context.Context,
	workers int,
	jobs []Job,
) []Result {
	jobCh := make(chan Job)
	resultCh := make(chan Result)

	var wg sync.WaitGroup

	// Worker function
	worker := func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case job, ok := <-jobCh:
				if !ok {
					return
				}

				err := processJob(ctx, job)

				select {
				case <-ctx.Done():
					return
				case resultCh <- Result{JobID: job.ID, Err: err}:
				}
			}
		}
	}

	// Start workers
	wg.Add(workers)
	for range workers {
		go worker()
	}

	// Feed jobs
	go func() {
		defer close(jobCh)
		for _, job := range jobs {
			select {
			case <-ctx.Done():
				return
			case jobCh <- job:
			}
		}
	}()

	// Close result channel after workers finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect results
	results := make([]Result, 0, len(jobs))
	for res := range resultCh {
		results = append(results, res)
	}

	return results
}

func processJob(ctx context.Context, job Job) error {
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(delay):
	}

	// 20% failure rate
	if rand.Intn(5) == 0 {
		return errors.New("job failed")
	}

	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	jobs := []Job{{1}, {2}, {3}, {4}, {5}}
	results := RunWorkerPool(ctx, 2, jobs)

	for _, r := range results {
		fmt.Println(r.JobID, r.Err)
	}
}
