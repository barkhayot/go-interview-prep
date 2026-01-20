# Context-aware Worker Pool

> Implement a worker pool that processes jobs concurrently.


### Requirements

- Each job:
    - Has an id (int)
    - Takes random time between 100–500 ms to complete
    - May fail randomly (20% chance)
- A job returns:
    - `nil` on success
    - an `error` on failure

- Worker Pool behavior:
    - Implement a worker pool with:
    - N workers (configurable)
    - A job queue
    - Graceful shutdown using context.Context
- Rules:
    - Workers should stop immediately when the context is cancelled
    - No goroutine leaks
    - All started jobs must either:
        - complete, or
        - be cancelled by context

### Interface

```go
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
) []Result
```
