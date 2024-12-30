package cache_test

import (
	"testing"
	"time"

	"github.com/thesayedirfan/weather/internal/cache"
	"github.com/thesayedirfan/weather/types"
)

func TestNewWeatherCache(t *testing.T) {
	weatherCache := cache.NewWeatherCache()
	if weatherCache == nil {
		t.Fatalf("expected non-nil WeatherCacheStore, got nil")
	}
}

func TestWeatherCache_SetAndGet_Success(t *testing.T) {
	weatherCache := cache.NewWeatherCache()

	key := "NewYork"
	value := types.Response{
		IP:       "192.168.1.1",
		Location: types.Location{City: "New York", Country: "USA"},
		Weather:  types.Weather{Temperature: "25", Humidity: "60", Description: "Clear Sky"},
	}

	weatherCache.Set(key, value, 5)

	cachedValue, found := weatherCache.Get(key)
	if !found {
		t.Fatalf("expected value to be found, got not found")
	}

	if cachedValue.IP != value.IP || cachedValue.Location.City != value.Location.City {
		t.Errorf("expected cached value %+v, got %+v", value, cachedValue)
	}
}

func TestWeatherCache_Get_Expired(t *testing.T) {
	weatherCache := cache.NewWeatherCache()

	key := "NewYork"
	value := types.Response{
		IP:       "192.168.1.1",
		Location: types.Location{City: "New York", Country: "USA"},
		Weather:  types.Weather{Temperature: "25", Humidity: "60", Description: "Clear Sky"},
	}

	weatherCache.Set(key, value, 1)
	time.Sleep(2 * time.Second)

	_, found := weatherCache.Get(key)
	if found {
		t.Fatalf("expected value to be expired, but it was found")
	}
}

func TestWeatherCache_Get_NonExistentKey(t *testing.T) {
	weatherCache := cache.NewWeatherCache()

	key := "NonExistentKey"

	_, found := weatherCache.Get(key)
	if found {
		t.Fatalf("expected value to not be found for non-existent key, but it was found")
	}
}
