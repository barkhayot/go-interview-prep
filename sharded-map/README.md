# Sharded Map (Concurrent Map)

> Implements a thread-safe map using sharding for improved concurrency.

### Problem

Design a map that supports concurrent access by splitting the data into multiple shards, each protected by its own lock. This reduces lock contention and improves performance for high-concurrency workloads.

### Requirements

- Split the map into N shards (configurable)
- Each shard is protected by its own mutex
- Get and Set operations use a hash function to select the shard
- Safe for concurrent use
- No external libraries (standard Go only)

### Interface

```go
type ShardedMap struct {
    // ...shards and fields...
}

func NewShardedMap(shards int) *ShardedMap
func (sm *ShardedMap) Get(key string) (any, bool)
func (sm *ShardedMap) Set(key string, val any)
```

### Example

```go
sm := NewShardedMap(4)
sm.Set("foo", "bar")
value, found := sm.Get("foo")
if found {
    fmt.Println("Found:", value)
}
```

---

See `main.go` for the implementation and runnable example.
