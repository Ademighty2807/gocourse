package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:  limit,
		window: window,
	}
}

func (r *RateLimiter) Allow() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	if now.After(r.resetTime) {
		r.count = 0
		r.resetTime = now.Add(r.window)
	}
	if r.count < r.limit {
		r.count++
		return true
	}
	return false
}

func main() {
	rateLimiter := NewRateLimiter(7, 3*time.Second)
	var wg sync.WaitGroup
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rateLimiter.Allow() {
				fmt.Println("Request allowed")
			} else {
				fmt.Println("Request denied")
			}
		}()
		// time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
