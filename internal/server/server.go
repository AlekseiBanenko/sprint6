package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	httpServer *http.Server
	Logger     *log.Logger
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		httpServer: srv,
		Logger:     logger,
	}
}

func (s *Server) Start() error {
	s.Logger.Println("Starting server on", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
