package main

import "sync"

type cache struct {
	mu    sync.RWMutex
	state map[string]any
}

type ShardedMap struct {
	shards    []*cache
	numShards int
}

func NewShardedMap(shards int) *ShardedMap {
	sm := &ShardedMap{
		shards:    make([]*cache, shards),
		numShards: shards,
	}

	for i := 0; i < sm.numShards; i++ {
		cache := cache{
			state: make(map[string]any),
		}
		sm.shards[i] = &cache
	}

	return sm
}

func (sm *ShardedMap) Get(key string) (any, bool) {
	hash := hashit(key)
	shardNum := hash % len(sm.shards)

	return sm.shards[shardNum].get(key)
}

func (sm *ShardedMap) Set(key string, val any) {
	hash := hashit(key)
	shardNum := hash % len(sm.shards)

	sm.shards[shardNum].set(key, val)
}

func (c *cache) get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.state[key]
	return v, ok
}

func (c *cache) set(key string, val any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.state[key] = val
}

func hashit(key string) int {
	return len(key) + 1
}

func main() {
	sm := NewShardedMap(4)

	sm.Set("foo", "bar")
	value, found := sm.Get("foo")

	if found {
		println("Found:", value.(string))
	} else {
		println("Not found")
	}
}
