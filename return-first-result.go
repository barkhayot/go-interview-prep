package main

import (
	"context"
	"fmt"
)

func Get(ctx context.Context, address string) (string, error) {
	return fmt.Sprintf("Got result from %s", address), nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	vals := []string{"res1", "res2", "res3"}
	res := make(chan string, 1) // buffer of 1 to avoid blocking

	for _, a := range vals {
		add := a
		go func() {
			val, err := Get(ctx, add)
			if err != nil {
				return
			}

			// send result or exit if cancelled
			select {
			case res <- val:
			case <-ctx.Done():
				return
			}
		}()
	}

	// wait for first result
	first := <-res
	fmt.Println(first)

	// cancel all other goroutines
	cancel()
}
