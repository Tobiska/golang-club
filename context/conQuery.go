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

func waitOne(servicesQuery ...func(context.Context)) {
	ch := make(chan struct{}, 0)

	ctx, cancel := context.WithCancel(context.Background())
	for _, query := range servicesQuery {
		go func() {
			servicesQuery()
			ch <- struct{}{}
		}
	}

	<-ch
	cancel()
}

func main() {
	waitOne(wrap(ctx))
}
