package main

import (
	"log"

	"github.com/elahe-dastan/simplesurance-assignment/internal/http"
)

func main() {
	srv := http.New()

	if err := srv.Run(); err != nil {
		log.Fatalf("running server failed %s", err)
	}
}
