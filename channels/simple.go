package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	i, ok := <-c
	fmt.Println(i, ok)

	c <- 10
	fmt.Println(<-c)

	c1 := make(chan int, 2)

	c1 <- 12

	fmt.Println(<-c1)
	fmt.Println(<-c1)
	fmt.Println("finish")
}
