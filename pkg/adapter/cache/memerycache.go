package cache

import (
	"sync"
	"time"
)

type Data struct {
	value    any
	expire   time.Duration
	createAt time.Time
}

func (d *Data) Expired() bool {
	return time.Now().Sub(d.createAt) > d.expire
}

type MemoryCache struct {
	mu    sync.RWMutex
	store map[string]Data
}

var memoryCache Cache

func init() {
	_memoryCache := NewMemoryCache()
	memoryCache = _memoryCache
}

// Clean implements Cache.
func (m *MemoryCache) Clean() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.store = map[string]Data{}
}

// Delete implements Cache.
func (m *MemoryCache) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.store, key)
	return nil
}

// Get implements Cache.
func (m *MemoryCache) Get(key string) (value any, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if v, ok := m.store[key]; ok {
		if !v.Expired() {
			return v.value, nil
		} else {
			delete(m.store, key)
		}
	}
	return nil, nil
}

// Set implements Cache.
func (m *MemoryCache) Set(key string, value any, expire time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data := Data{
		value:    value,
		expire:   expire,
		createAt: time.Now(),
	}
	m.store[key] = data
	return nil
}

func (m *MemoryCache) Gc() {
	for {
		select {
		case <-time.After(1 * time.Second):
			keysToExpire := []string{}
			m.mu.RLock()
			for k, v := range m.store {
				if v.Expired() {
					keysToExpire = append(keysToExpire, k)
				}
			}
			m.mu.Unlock()

			for _, v := range keysToExpire {
				m.Delete(v)
			}
		}
	}
}

func NewMemoryCache() Cache {
	return &MemoryCache{
		mu:    sync.RWMutex{},
		store: map[string]Data{},
	}
}

func GetMemoryCache() Cache {
	return memoryCache
}
