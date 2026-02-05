package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
}

func NewRateLimiter(rate int, per time.Duration) *RateLimiter {
	r := &RateLimiter{
		tokens: make(chan struct{}, rate),
	}

	// fill it on init
	for range rate {
		r.tokens <- struct{}{}
	}

	// launch filler of bucket
	interval := per / time.Duration(rate)
	ticker := time.NewTimer(interval)

	go func() {
		for range ticker.C {
			select {
			// fill channel on needed interval
			case r.tokens <- struct{}{}:
			// do nothing when channel is full
			default:
			}
		}
	}()

	return r
}

func (r *RateLimiter) Allow() bool {
	select {
	case <-r.tokens:
		return true
	default:
		return false
	}
}

func main() {
	r := NewRateLimiter(5, time.Second)

	requests := 10

	for range requests {
		allowed := r.Allow()
		if !allowed {
			fmt.Println("request is not allwed")
		}

		fmt.Println("request is allowed")
	}
}
