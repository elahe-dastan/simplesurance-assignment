package http

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func New() Server {
	srv := &http.Server{
		Addr: ":1378",
	}

	return Server{
		srv: srv,
	}
}

func (s Server) Register(h http.Handler) {
	s.srv.Handler = h
}

func (s Server) Run() error {
	log.Printf("listening on %s\n", s.srv.Addr)

	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("serving and listening failed %w", err)
	}

	return nil
}
