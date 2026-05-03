package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Port uint16
}

func New(port string) (http.Server, error) {
	s := Server{}
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

	return r
}

func (s *Server) homeRoute(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("hello"))
	return nil
}
