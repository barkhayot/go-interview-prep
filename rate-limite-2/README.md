# Rate Limiter (Simple Token Bucket)

> Implements a basic rate limiter using the token bucket algorithm.

### Problem

Design a rate limiter that allows a fixed number of requests per time interval. Each request consumes a token from a bucket, and tokens are refilled at a steady rate. If no tokens are available, the request is denied.

### Requirements

- Fixed capacity (max tokens in the bucket)
- Refill tokens at a constant rate
- Each request consumes 1 token
- Requests are allowed only if a token is available
- Thread-safe and suitable for concurrent use
- No external libraries (standard Go only)

### Interface

```go
type RateLimiter struct {
    tokens chan struct{}
}

func NewRateLimiter(rate int, per time.Duration) *RateLimiter
func (r *RateLimiter) Allow() bool
```

### Example

```go
r := NewRateLimiter(5, time.Second)

for i := 0; i < 10; i++ {
    if r.Allow() {
        fmt.Println("Request", i, "allowed")
    } else {
        fmt.Println("Request", i, "denied")
    }
}
```

---

See `main.go` for the implementation and runnable example.
