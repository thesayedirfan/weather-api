package cache

import (
	"sync"
	"time"

	"github.com/thesayedirfan/weather/types"
)

type WeatherCache struct {
	Value     *types.Response
	ExpiresAt time.Time
}

type WeatherCacheStore struct {
	cache map[string]WeatherCache
	mu    sync.Mutex
}

func NewWeatherCache() *WeatherCacheStore {
	return &WeatherCacheStore{
		cache: make(map[string]WeatherCache),
	}
}

func (c *WeatherCacheStore) Set(key string, value types.Response, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = WeatherCache{
		Value:     &value,
		ExpiresAt: time.Now().Add(time.Duration(ttl) * time.Second),
	}
}

func (c *WeatherCacheStore) Get(key string) (*types.Response, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, found := c.cache[key]
	if !found || time.Now().After(item.ExpiresAt) {
		return nil, false
	}
	return item.Value, true
}
