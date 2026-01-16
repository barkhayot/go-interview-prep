package main

import "sync"

type PubSub struct {
	subscribers map[string][]chan string
	mu          sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)
	return ch
}

func (ps *PubSub) Publish(topic, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, ch := range ps.subscribers[topic] {
		ch <- message
	}
}

func (ps *PubSub) Unsubscribe(topic string, ch <-chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subs := ps.subscribers[topic]
	for i, subscriber := range subs {
		if subscriber == ch {
			ps.subscribers[topic] = append(subs[:i], subs[i+1:]...)
			close(subscriber)
			break
		}
	}
}

func main() {
	ps := NewPubSub()

	subscriber := ps.Subscribe("news")
	ps.Publish("news", "Breaking News!")

	message := <-subscriber
	println("Received:", message)

	ps.Unsubscribe("news", subscriber)
}
