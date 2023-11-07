package cache

import "time"

type Option struct {
}

type Cache interface {
	Get(key string) (value any, err error)
	Set(key string, value any, expire time.Duration) error
	Delete(key string) error
	Clean()
}
