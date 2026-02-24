package main

import "sync/atomic"

type Once struct {
	done atomic.Bool
}

func (o *Once) Do(f func()) {
	if o.done.Load() {
		return
	}

	if o.done.CompareAndSwap(false, true) {
		f()
	}
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
