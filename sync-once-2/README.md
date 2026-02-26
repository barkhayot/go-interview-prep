# Sync-once-2 (Atomic sync.Once)

> Implements a version of sync.Once using atomic operations for one-time initialization in concurrent programs.

### Problem

Design a synchronization primitive that ensures a function is executed only once, even if called from multiple goroutines, using atomic operations.

### Requirements

- Only one execution of the function, even with concurrent calls
- Safe for concurrent use
- No external libraries (standard Go only)

### Interface

```go
type Once struct {}

func (o *Once) Do(f func())
```

### Example

```go
var once Once

once.Do(func() {
  println("This will only be printed once.")
})

once.Do(func() {
  println("This will not be printed.")
})
```

---
