package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens       chan struct{}
	refillTicker *time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:       make(chan struct{}, rateLimit),
		refillTicker: &refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(*rl.refillTicker)
	go func() {
		for range ticker.C {
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}()
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {
	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
