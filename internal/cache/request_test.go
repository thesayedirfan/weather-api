package cache

import (
	"testing"
)

func TestNewRequestCache(t *testing.T) {
	requestCache := NewRequestCache()
	if requestCache == nil {
		t.Fatalf("expected non-nil RequestCache, got nil")
	}
	if len(requestCache.cache) != 0 {
		t.Errorf("expected empty cache, got length %d", len(requestCache.cache))
	}
}

func TestRequestCache_Increment(t *testing.T) {
	requestCache := NewRequestCache()
	requestCache.Increment("192.168.1.1")

	if requestCache.Get("192.168.1.1") != 1 {
		t.Errorf("expected count 1 after increment, got %d", requestCache.Get("192.168.1.1"))
	}

	requestCache.Increment("192.168.1.1")
	if requestCache.Get("192.168.1.1") != 2 {
		t.Errorf("expected count 2 after second increment, got %d", requestCache.Get("192.168.1.1"))
	}
}

func TestRequestCache_Decrement(t *testing.T) {
	requestCache := NewRequestCache()
	requestCache.Increment("192.168.1.1")
	requestCache.Increment("192.168.1.1")

	requestCache.Decrement("192.168.1.1")
	if requestCache.Get("192.168.1.1") != 1 {
		t.Errorf("expected count 1 after decrement, got %d", requestCache.Get("192.168.1.1"))
	}

	requestCache.Decrement("192.168.1.1")
	if requestCache.Get("192.168.1.1") != 0 {
		t.Errorf("expected count 0 after second decrement, got %d", requestCache.Get("192.168.1.1"))
	}
}
