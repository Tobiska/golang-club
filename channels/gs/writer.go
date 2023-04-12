package main

import (
	"context"
	"fmt"
	"time"
)

type WriterStdout struct {
	err chan error
}

func NewWriter() *WriterStdout {
	return &WriterStdout{
		err: make(chan error),
	}
}

func (w *WriterStdout) Run() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("write")
	}
}

func (w *WriterStdout) Notify() chan error {
	return w.err
}

func (w *WriterStdout) Shutdown(_ context.Context) error {
	time.Sleep(1 * time.Second)
	fmt.Println("writer finished")
	return nil
}
