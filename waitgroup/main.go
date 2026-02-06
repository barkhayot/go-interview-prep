package main

import (
	"fmt"
	"sync"
	"time"
)

type WaitGroup struct {
	count int
	done  chan struct{}
	mu    sync.Mutex
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

func (w *WaitGroup) Add(n int) {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.count == 0 && n > 0 {
		// reset for reuse
		w.done = make(chan struct{})
	}

	w.count += n
	if w.count < 0 {
		panic("negative WaitGroup counter")
	}
}

func (w *WaitGroup) Done() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.count--
	if w.count < 0 {
		panic("negative WaitGroup counter")
	}

	if w.count == 0 {
		close(w.done)
	}
}

func (w *WaitGroup) Wait() {
	<-w.done
}

func main() {
	wg := NewWaitGroup()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("executed!")
		time.Sleep(time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("executed! one more time")
		time.Sleep(time.Second)
	}()

	wg.Wait()
	fmt.Println("Try programiz.pro")
}
