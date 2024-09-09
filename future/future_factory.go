package main

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

const (
	defaultCountRetries       = 3
	defaultTimeBetweenRetries = 200 * time.Millisecond
)

type Future func(ctx context.Context) <-chan Data

type Data struct {
	value string
	Error error
}

type Option func(f *FutureFactory)

type FutureFactory struct {
	countRetries       int
	timeBetweenRetries time.Duration
	beforeFunc         func(ctx context.Context) error
	afterFunc          func(ctx context.Context) error
	group              *errgroup.Group
}

func NewFactory(ctx context.Context, options ...Option) *FutureFactory {
	ff := &FutureFactory{
		countRetries:       defaultCountRetries,
		timeBetweenRetries: defaultTimeBetweenRetries,
	}

	for _, opt := range options {
		opt(ff)
	}

	ff.group, _ = errgroup.WithContext(ctx)

	return ff
}

func (f *FutureFactory) MakeFuture(deferFunc func(ctx context.Context) Data) Future {
	return func(ctx context.Context) <-chan Data {
		outCh := make(chan Data, 1)

		f.group.Go(func() error {
			var result Data

			if err := f.beforeFunc(ctx); err != nil {
				return err
			}

			for i := 0; i < f.countRetries; i++ {
				result = deferFunc(ctx)
				if result.Error == nil {
					break
				}
				time.Sleep(f.timeBetweenRetries)
			}

			outCh <- result
			close(outCh)

			if err := f.beforeFunc(ctx); err != nil {
				return err
			}

			return nil
		})

		return outCh
	}
}

func WithCountRetries(retries int) Option {
	return func(f *FutureFactory) {
		f.countRetries = retries
	}
}

func WithBeforeFunc(before func(ctx context.Context) error) Option {
	return func(f *FutureFactory) {
		f.beforeFunc = before
	}
}

func WithAfterFunc(after func(ctx context.Context) error) Option {
	return func(f *FutureFactory) {
		f.afterFunc = after
	}
}

func WithTimeBetweenRetries(retries int) Option {
	return func(f *FutureFactory) {
		f.countRetries = retries
	}
}

func main() {

}
