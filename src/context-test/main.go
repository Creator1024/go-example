package main

import (
	"context"
	"fmt"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func CancelTest() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
			}
			time.Sleep(time.Millisecond * 10)
		}(i, ctx)
		fmt.Println(i, "is Cancelled.")
	}
	cancel()
	time.Sleep(time.Second * 2)
}

func main() {
	CancelTest()
}
