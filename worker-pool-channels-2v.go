// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import (
    "fmt"
    "time"
    "context"
    "sync"
)

const (
    workers = 5
    timeout = 100 * time.Millisecond
)

func process(n int) int {
    return n*2
}

func main() {
  fmt.Println("Try programiz.pro")
  
  ctx, cancel := context.WithTimeout(context.Background(), timeout)
  defer cancel()
  
  jobs := []int{1,2,3,4,5,6,7}
  c := make(chan int)
  r := make(chan int)
  w := sync.WaitGroup{}

  for i:=0; i < workers; i++ {
      w.Add(1)
      go func(ctx context.Context, c, r chan int, w *sync.WaitGroup){
          defer w.Done()
          
          for {
              select {
                  case j, ok := <- c:
                    if !ok{
                        return
                    }
                    v := process(j)
                    r <- v
                  case <-ctx.Done():
                    return
              }
          }
      }(ctx, c, r, &w)
  }
  
  go func(){
      for _, j := range jobs {
          c <-j
      }
      
      close(c)
  }()
  
  go func(){
     defer close(r)
     w.Wait()
  }()
  
  for res := range r {
      fmt.Println(res)
  }
  
}





