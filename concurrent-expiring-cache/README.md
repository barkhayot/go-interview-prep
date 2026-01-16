# New Problem — Concurrent Expiring Cache (TTL Cache)

Design an in-memory cache with TTL.

### Requirements

- Each key has a value and an expiration time (TTL).
- Get returns value only if not expired.
- Expired keys must not be returned.
- Cache must be safe for concurrent access.
- You do not need to proactively clean expired keys (lazy cleanup is fine).

### Interface

```go
type Cache struct {}

func NewCache() *Cache
func (c *Cache) Set(key string, value string, ttl time.Duration)
func (c *Cache) Get(key string) (string, bool)
```

### Example

```go
c := NewCache()

c.Set("a", "1", time.Second)

time.Sleep(500 * time.Millisecond)
c.Get("a") // ("1", true)

time.Sleep(600 * time.Millisecond)
c.Get("a") // ("", false)
```
