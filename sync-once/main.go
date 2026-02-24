package main

import "sync"

type Once struct {
	done bool
	mu   sync.Mutex
}

func (o *Once) Do(f func()) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if o.done {
		return
	}

	f()
	o.done = true
}

func main() {
	var once Once

	once.Do(func() {
		println("This will only be printed once.")
	})

	once.Do(func() {
		println("This will not be printed.")
	})
}
