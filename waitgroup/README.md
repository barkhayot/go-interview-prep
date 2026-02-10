# Custom WaitGroup (Synchronization Primitive)

> Implements a simplified version of Go's sync.WaitGroup for goroutine synchronization.

### Problem

Design a synchronization primitive that allows one or more goroutines to wait until a set of operations being performed in other goroutines completes. This is a reimplementation of Go's standard `sync.WaitGroup`.

### Requirements

- Support Add, Done, and Wait methods
- Wait blocks until the counter is zero
- Add increments the counter, Done decrements it
- Panic if the counter becomes negative
- Safe for concurrent use
- No external libraries (standard Go only)

### Interface

```go
type WaitGroup struct {
    // ...fields...
}

func NewWaitGroup() *WaitGroup
func (w *WaitGroup) Add(n int)
func (w *WaitGroup) Done()
func (w *WaitGroup) Wait()
```

### Example

```go
wg := NewWaitGroup()
wg.Add(2)
go func() {
    defer wg.Done()
    // ...work...
}()
go func() {
    defer wg.Done()
    // ...work...
}()
wg.Wait()
fmt.Println("All done!")
```

---

See `main.go` for the implementation and runnable example.
