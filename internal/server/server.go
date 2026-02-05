package server

import (
	"log"
	"net/http"
	"time"

	"handlers"
)

type Server struct {
	logger *log.Logger
	server *http.Server
}

func New(l *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/upload", handlers.Upload)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: l,
		server: srv,
	}
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
