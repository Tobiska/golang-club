package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"os/signal"
	"syscall"
	"time"
)

const (
	listenAddr      = "127.0.0.1:8080"
	shutdownTimeout = 5 * time.Second
)

func main() {
	notifyCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	c := &Closer{}
	//srv
	m := chi.NewRouter()
	handler := NewHandler()
	handler.Register(m)

	srv := NewSrv(listenAddr, m)
	c.Add(srv.Shutdown)
	go srv.Run()

	//writer
	writer := NewWriter()
	c.Add(writer.Shutdown)
	go writer.Run()

	select {
	case <-srv.Notify():
	case <-writer.Notify():
	case <-notifyCtx.Done():
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := c.Close(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}
