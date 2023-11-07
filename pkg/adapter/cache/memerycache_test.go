package cache

import (
	"sync"
	"testing"
	"time"
)

func TestMemoryCache_Clean(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		store map[string]Data
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
		{
			name: "clean cache",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:  "abc",
						expire: 1 * time.Second,
					},
					"b": {
						value:  333,
						expire: 10 * time.Second,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				mu:    tt.fields.mu,
				store: tt.fields.store,
			}
			m.Clean()
			a, err := m.Get("a")
			if err != nil {
				t.Errorf("get a err: %v\n", err)
				return
			}
			t.Logf("a: %+v\n", a)
			b, err := m.Get("b")
			if err != nil {
				t.Errorf("get b err: %v\n", err)
				return
			}
			t.Logf("a: %+v\n", b)
		})
	}
}

func TestMemoryCache_Delete(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		store map[string]Data
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "delete cache1",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:  "abc",
						expire: 1 * time.Second,
					},
				},
			},
		},
		{
			name: "delete cache2",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:  "abc",
						expire: 10 * time.Second,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				mu:    tt.fields.mu,
				store: tt.fields.store,
			}
			if e := m.Delete("a"); e != nil {
				t.Error(e)
			}
			a, err := m.Get("a")
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("a: %+v", a)

		})
	}
}

func TestMemoryCache_Get(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		store map[string]Data
	}
	type args struct {
		key   string
		sleep int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue any
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name: "get cache1",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:    "abc",
						expire:   1 * time.Second,
						createAt: time.Now(),
					},
				},
			},
			args: args{
				key:   "a",
				sleep: 0,
			},
		},
		{
			name: "get cache2",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:    "abcdefg",
						expire:   3 * time.Second,
						createAt: time.Now(),
					},
				},
			},
			args: args{
				key:   "a",
				sleep: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				mu:    tt.fields.mu,
				store: tt.fields.store,
			}
			time.Sleep(time.Duration(tt.args.sleep) * time.Second)
			gotValue, err := m.Get(tt.args.key)
			if err != nil {
				t.Error(err)
				return
			}
			t.Logf("gotValue: %+v\n", gotValue)
		})
	}
}

func TestMemoryCache_Set(t *testing.T) {
	type fields struct {
		mu    sync.RWMutex
		store map[string]Data
	}
	type args struct {
		key    string
		value  any
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "set cache1",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:  "abc",
						expire: 1 * time.Second,
					},
				},
			},
		},
		{
			name: "set cache2",
			fields: fields{
				mu: sync.RWMutex{},
				store: map[string]Data{
					"a": {
						value:  "abc",
						expire: 10 * time.Second,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemoryCache{
				mu:    tt.fields.mu,
				store: tt.fields.store,
			}
			if err := m.Set(tt.args.key, tt.args.value, tt.args.expire); err != nil {
				t.Errorf("MemoryCache.Set() error = %v\n", err)
			}
		})
	}
}
