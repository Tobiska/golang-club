package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func longRequest() {
	time.Sleep(2 * time.Second)
	http.Get("https://www.google.ru/")
}

func wrap(ctx context.Context) {
	ch := make(chan struct{})
	go func() {
		longRequest()
		ch <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("done")
		return
	case <-ch:
		fmt.Println("request finish")
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wrap(ctx)
}
