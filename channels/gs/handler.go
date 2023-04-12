package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (h *Handler) Register(m *chi.Mux) {
	m.Get("/", h.Handle)
}
