package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type server struct {
	port uint16
}

func New(port string) (http.Server, error) {
	s := server{}
	portInt, err := fmt.Sscanf(port, "%d", &s.port)
	if err != nil {
		return http.Server{}, err
	}
	if portInt != 1 {
		return http.Server{}, fmt.Errorf("invalid port")
	}

	return http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.routes(),
	}, nil
}

func (s *server) routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", s.homeRoute)

	return r
}

func (s *server) homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
