# Concurrent Pub/Sub (In-Memory)

> Design a simple publish–subscribe system.

### Requirements

- Multiple subscribers can subscribe to a topic.
- A publisher can publish messages to a topic.
- Each subscriber receives all messages published after it subscribes.
- Publishing must not block on slow subscribers.
- Must be safe for concurrent use.

### Interface

```go
type PubSub struct {}

func NewPubSub() *PubSub

func (ps *PubSub) Subscribe(topic string) <-chan string
func (ps *PubSub) Publish(topic string, msg string)
func (ps *PubSub) Unsubscribe(topic string, ch <-chan string)

```

### Example

```go
ps := NewPubSub()

ch1 := ps.Subscribe("news")
ch2 := ps.Subscribe("news")

ps.Publish("news", "hello")

fmt.Println(<-ch1) // "hello"
fmt.Println(<-ch2) // "hello"

```
