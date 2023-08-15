package main

import (
	"context"
	"fmt"
	"time"
)

func mergeContext(ctx, cancelCtx context.Context) (context.Context, func()) {
	ctx, cancel := context.WithCancelCause(ctx)
	stop := context.AfterFunc(cancelCtx, func() {
		cancel(context.Cause(cancelCtx))
	})

	return ctx, func() {
		stop()
		cancel(context.Canceled)
	}
}

func main() {
	parentCtx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	childCtx := context.WithValue(parentCtx, 1, 2)

	childCtx = context.WithoutCancel(childCtx)

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1 second")

	case <-childCtx.Done():
		fmt.Println("child ctx done")
	}

	ctx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Millisecond, fmt.Errorf("bebra"))
	<-ctx.Done()

	fmt.Println(ctx.Err(), context.Cause(ctx))

	//mergeCtx

}
