package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/elahe-dastan/simplesurance-assignment/internal"
)

type Server struct {
	srv        *http.Server
	hitCounter *internal.HitCounter
}

func New(hitCounter *internal.HitCounter) Server {
	srv := &http.Server{
		Addr: ":1378",
	}

	return Server{
		srv:        srv,
		hitCounter: hitCounter,
	}
}

func (s Server) Register(h http.Handler) {
	s.srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// New logic to handle hits
		timestamp := time.Now().Unix()
		s.hitCounter.Hit(timestamp)

		hitCount := s.hitCounter.GetHitCount()

		_, err := fmt.Fprintf(w, "Total Hits: %d", hitCount)
		if err != nil {
			return
		}
	})
}

func (s Server) Run() error {
	if err := s.srv.ListenAndServe(); err != nil {
		return fmt.Errorf("serving and listening failed %w", err)
	}

	return nil
}
