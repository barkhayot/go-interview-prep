package main

import (
	"sync"
	"time"
)

type bucket struct {
	tokens     float64
	lastRefill time.Time
	refillRate float64 // per-user
}

type RateLimiter struct {
	capacity   float64
	refillRate float64 // tokens per second
	users      map[string]*bucket
	mu         sync.Mutex
}

func NewRateLimiter(capacity, refillRate float64) *RateLimiter {
	return &RateLimiter{
		capacity:   capacity,
		refillRate: refillRate,
		users:      make(map[string]*bucket),
	}

}

func (rl *RateLimiter) Allow(userID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	b, exists := rl.users[userID]
	if !exists {
		b = &bucket{
			tokens:     rl.capacity,
			lastRefill: now,
			refillRate: rl.refillRate,
		}
		rl.users[userID] = b
	}

	elapsed := now.Sub(b.lastRefill).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > rl.capacity {
		b.tokens = rl.capacity
	}
	b.lastRefill = now
	if b.tokens >= 1 {
		b.tokens -= 1
		return true
	}
	return false
}
