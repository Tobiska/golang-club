package main

import (
	"context"
	"errors"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("some error")
	})

	g.Go(func() error {
		<-ctx.Done()
		return errors.New("Some error")
	})

	println(g.Wait())
}
