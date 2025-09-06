package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	worker = 3
	interval = time.Millisecond * 200
)

func process(id, job int,){
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("Worker %d processed job %d\n", id, job)
}

func workerHandler(ctx context.Context, workerId int, channel chan int){
	ticker := time.NewTicker(interval)
	
	for {
    select {
    case job := <-channel:
        <-ticker.C
        process(workerId, job,)
    case <-ctx.Done():
        return
    }
}
}

func main() {
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  channel := make(chan int)
  jobs := []int{1,2,3,4,5,6,7,8,9,10}

  var wg sync.WaitGroup
  wg.Add(worker)
  for i := 1; i <= worker; i++ {
    go func(id int) {
      defer wg.Done()
      workerHandler(ctx, id, channel)
    }(i)
  }

  go func() {
    for _, job := range jobs {
      channel <- job
    }
    close(channel)
  }()

  wg.Wait()
}
