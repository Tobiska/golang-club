package main

import "time"

func main() {
	ch := make(chan int32, 4)

	ch <- 1
	ch <- 2
	_ = <-ch
	ch <- 3
	_ = <-ch
	//
	go func() {
		for {
			ch <- 42
		}
	}()

	time.Sleep(1 * time.Second)

	close(ch)
	close(ch)
	return
}
