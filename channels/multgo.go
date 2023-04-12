package main

import "fmt"

func main() {
	cnt := 10
	quit := make(chan struct{}, cnt)
	for i := 0; i < cnt; i++ {
		go func() {
			fmt.Println(i)
			quit <- struct{}{}
		}()
	}

	for i := 0; i < cnt; i++ {
		<-quit
	}
}
