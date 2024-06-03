package server

import (
	"github.com/kodaf1/gometrics/internal/repositories"
	"net/http"
)

type Server struct {
	Storage repositories.StorageRepository
}

func NewServer(storage repositories.StorageRepository) *Server {
	return &Server{
		Storage: storage,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/update/", s.UpdateMetric)
	return mux
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: s.Handler(),
	}

	return httpServer.ListenAndServe()
}
