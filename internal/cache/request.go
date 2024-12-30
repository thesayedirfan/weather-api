package cache

import "sync"

type Request struct {
	cache map[string]int
	mu    sync.Mutex
}

func NewRequestCache() *Request {
	return &Request{
		cache: make(map[string]int),
	}
}

func (c *Request) Get(ip string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.cache[ip]
}

func (c *Request) Increment(ip string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[ip]++
}

func (c *Request) Decrement(ip string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[ip]--
}
