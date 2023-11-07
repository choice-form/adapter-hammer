package cache

import (
	"reflect"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRedisClient_Set(t *testing.T) {
	cli := NewRedisClient(&Config{
		Host:     "localhost",
		Port:     0,
		Username: "",
		Password: "",
	})
	err := cli.Set("abc", 123, 0)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(3 * time.Second)
	result, err := cli.Get("abc")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("result", result)
}

func TestRedisClient_Get(t *testing.T) {
	type fields struct {
		config *Config
		client *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue any
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &RedisClient{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			gotValue, err := rdb.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("RedisClient.Get() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
