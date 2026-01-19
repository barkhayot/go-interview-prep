# Design & Implement: Rate Limiter (API Gateway style)

> Design a rate limiter that limits requests per user using the Token Bucket algorithm.

### Requirements

- Each user has:
  - capacity (max tokens)
  - refill_rate (tokens per second)
- Every request:
  - Consumes 1 token
  - Is allowed only if a token is available
- The rate limiter must:
  - Be thread-safe
  - Work in high-concurrency
  - Support multiple users
- No external libraries (standard Go only)

### Interface

```go
type RateLimiter struct {
    // your fields
}

func NewRateLimiter(capacity int, refillRate float64) *RateLimiter

// Allow returns true if request is allowed for this user
func (rl *RateLimiter) Allow(userID string) bool


```

### Example

```go
rl := NewRateLimiter(5, 1.0) // capacity=5, refillRate=1 token/sec
userID := "user1"
for i := 0; i < 10; i++ {
    if rl.Allow(userID) {
        fmt.Println("Request", i, "allowed")
    } else {
        fmt.Println("Request", i, "denied")
    }
    time.Sleep(200 * time.Millisecond) // simulate time between requests
}


```
