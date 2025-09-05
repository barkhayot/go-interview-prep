package main

import (
	"fmt"
	"sync"
)

func fanIn(channels ...<-chan int) <-chan int {
  out := make(chan int)  
	var wg sync.WaitGroup

	run := func(ch <-chan int) {
		defer wg.Done()
		for v1 := range ch{
			out <- v1
		}
	}

	wg.Add(len(channels))
	for _, channel := range channels {
		go run(channel)
	}

	go func() {
		wg.Wait()
    close(out)
	}()

	return out
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	out := fanIn(chan1, chan2)

	go func(){
		for i := 0; i <= 100; i++ {
			chan1 <- i
		}
		close(chan1)
	}()

	go func(){
    for i := 0; i <= 100; i++ {
      chan2 <- i
    }
    close(chan2)
  }()

	for v := range out {
    fmt.Println(v)
}
}
