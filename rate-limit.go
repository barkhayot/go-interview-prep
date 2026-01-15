package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu     sync.Mutex
	window time.Duration
	limit  int
	users  map[string]*userBucket
}

type userBucket struct {
	times []time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		window: window,
		limit:  limit,
		users:  make(map[string]*userBucket),
	}
}

func (r *RateLimiter) Allow(userID string) bool {
	now := time.Now()

	r.mu.Lock()
	defer r.mu.Unlock()

	b, ok := r.users[userID]
	if !ok {
		b = &userBucket{}
		r.users[userID] = b
	}

	// remove expired timestamps
	validFrom := now.Add(-r.window)
	idx := 0
	for idx < len(b.times) && b.times[idx].Before(validFrom) {
		idx++
	}
	b.times = b.times[idx:]

	if len(b.times) >= r.limit {
		return false
	}

	b.times = append(b.times, now)
	return true
}

func main() {
	r := NewRateLimiter(10, 1*time.Second)
	times := 10000

	go func() {
		for range times {
			time.Sleep(500 * time.Microsecond)
			allowed := r.Allow("user")
			if allowed {
				fmt.Println("allowed")
			} else {
				fmt.Println("not Allowed")
			}

		}
	}()

	time.Sleep(10 * time.Second)
}
