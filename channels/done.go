package main

import (
	"context"
	"fmt"
	"time"
)

func read(ctx context.Context, ch chan int) {
	for {
		if err := ctx.Err(); err != nil {
			return
		}
		select {
		case <-ctx.Done():
			return
		case t := <-ch:
			fmt.Printf("read %d\n", t)
		}
	}
}

func write(ctx context.Context, ch chan int) {
	for i := 0; i < 10000000000; i++ {
		if err := ctx.Err(); err != nil {
			return
		}
		select {
		case <-ctx.Done():
			close(ch)
			return
		case ch <- i:
			fmt.Printf("write %d\n", i)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int, 10)
	go read(ctx, ch)

	go write(ctx, ch)

	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("CANCEL")
	time.Sleep(2 * time.Second)
}
