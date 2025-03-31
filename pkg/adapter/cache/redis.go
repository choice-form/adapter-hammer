package cache

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
}

type RedisClient struct {
	config *Config
	client *redis.Client
}

// Clean implements Cache.
func (*RedisClient) Clean() {
	panic("unimplemented")
}

// Delete implements Cache.
func (rdb *RedisClient) Delete(key string) error {
	_, err := rdb.client.Del(context.Background(), key).Result()
	if err != nil {
		return err
	}
	return nil
}

// Get implements Cache.
func (rdb *RedisClient) Get(key string) (value any, err error) {
	return rdb.client.Get(context.Background(), key).Result()
}

// Set implements Cache.
func (rdb *RedisClient) Set(key string, value any, expire time.Duration) error {
	return rdb.client.Set(context.Background(), key, value, expire).Err()
}

func (rdb *RedisClient) Client() *redis.Client {
	return rdb.client
}

// func init() {
// 	cfg := config.GetConfig()
// 	rds := NewRedisClient(&cfg.Redis)
// 	redisClient = rds
// }

func NewRedisClient(cfg *Config) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       0,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	slog.Info("new redis client init", slog.String("ping", _ping))

	return &RedisClient{
		config: cfg,
		client: rdb,
	}
}

// func GetRedisClient() Cache {
// 	return redisClient
// }
