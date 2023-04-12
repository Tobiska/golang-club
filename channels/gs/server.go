package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

type HttpServer struct {
	r     *chi.Mux
	errCh chan error
	addr  string
}

func NewSrv(addr string, r *chi.Mux) *HttpServer {
	return &HttpServer{
		r:     r,
		addr:  addr,
		errCh: make(chan error),
	}
}

func (s *HttpServer) Notify() chan error {
	return s.errCh
}

func (s *HttpServer) Shutdown(_ context.Context) error {
	time.Sleep(3 * time.Second)
	fmt.Println("server finished")
	return nil
}

func (s *HttpServer) Run() {
	s.errCh <- http.ListenAndServe(s.addr, s.r)
}
