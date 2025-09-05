package main

import (
	"fmt"
	"time"
)
 

func main() {
	tunel := make(chan int)
	fmt.Println("launching app")
	// producer
	go func() {
		for i:=1; i<=100; i++{
			tunel <- i
		}
		close(tunel)
	}()

	// receiver
	for number := range tunel {
    fmt.Printf("received number is %d\n", number*number) // square it
  }

	time.Sleep(time.Minute)
}
