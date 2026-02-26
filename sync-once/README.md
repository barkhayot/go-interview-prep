# Sync-once (Custom sync.Once)

> Implements a simplified version of Go's sync.Once for one-time initialization in concurrent programs.

### Problem

Design a synchronization primitive that ensures a function is executed only once, even if called from multiple goroutines.

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
