// 限流
package ratelimiter

import (
	"context"

	"github.com/choice-form/adapter-hammer/pkg/logger"
	"golang.org/x/time/rate"
)

var limit *rate.Limiter

// 设置全局限流
func SetLimiter(r rate.Limit, b int) {
	limit = rate.NewLimiter(r, b)
}

func RateLimiter(ctx context.Context, fn func() error) error {
	err := limit.Wait(ctx)
	if err != nil {
		logger.Error("RateLimiter error", logger.Any("error", err))
		return err
	}
	return fn()
}
