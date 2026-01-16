package main

import (
	"sync"
	"time"
)

type entry struct {
	value    string
	expireAt time.Time
}

type Cache struct {
	mu sync.Mutex
	m  map[string]entry
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]entry),
	}
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[key] = entry{
		value:    value,
		expireAt: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e, ok := c.m[key]
	if !ok {
		return "", false
	}

	if time.Now().After(e.expireAt) {
		delete(c.m, key)
		return "", false
	}

	return e.value, true
}

func main() {
	cache := NewCache()

	cache.Set("foo", "bar", 1*time.Second)
	value, found := cache.Get("foo")
	if found {
		println("Found:", value)
	} else {
		println("Not found")
	}

	time.Sleep(3 * time.Second)
	value, found = cache.Get("foo")
	if found {
		println("Found:", value)
	} else {
		println("Not found")
	}
}
