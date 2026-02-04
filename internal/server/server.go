package server

import (
	"log"
	"net/http"
	"time"

	"sprint6/internal/handlers"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(logger *log.Logger) *Server {
	r := chi.NewRouter()
	r.Get("/", handlers.Index)
	r.Post("/upload", handlers.Upload)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
		ErrorLog:     logger,
	}

	return &Server{logger: logger, server: srv}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
