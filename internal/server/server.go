package server

import (
	"fmt"
	"net/http"

	"github.com/alcb1310/bca-proto/internal/database"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port uint16
	DB   database.Service
}

func New(port string, db database.Service) (http.Server, error) {
	s := Server{
		DB: db,
	}

	portInt, err := fmt.Sscanf(port, "%d", &s.Port)
	if err != nil {
		return http.Server{}, err
	}
	if portInt != 1 {
		return http.Server{}, fmt.Errorf("invalid port")
	}

	return http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Router(),
	}, nil
}

func (s *Server) Router() http.Handler {
	r := chi.NewRouter()

	r.Get("/", HandleErrors(s.homeRoute))
	r.Get("/health", HandleErrors(s.healthRoute))

	return r
}

func (s *Server) homeRoute(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello"))
	return nil
}

func (s *Server) healthRoute(w http.ResponseWriter, r *http.Request) error {
	err := s.DB.HealthCheck()
	if err != nil {
		return err
	}

	w.Write([]byte("server healthy"))
	return nil
}
