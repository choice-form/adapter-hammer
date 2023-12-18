package ratelimiter

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestRateLimiter(t *testing.T) {
	var mux = sync.WaitGroup{}
	t.Run("test limit", func(t *testing.T) {
		for k := 0; k < 20; k++ {
			mux.Add(1)
			go RateLimiter(context.Background(), func() error {
				fmt.Printf("k: %v\n", k)
				mux.Done()
				return nil
			})
		}

		mux.Wait()
	})
}
