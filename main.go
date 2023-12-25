package main

import (
	"log"

	"github.com/elahe-dastan/simplesurance-assignment/internal"
	"github.com/elahe-dastan/simplesurance-assignment/internal/http"
)

func main() {
	hitCounter := &internal.HitCounter{}
	srv := http.New(hitCounter)

	if err := srv.Run(); err != nil {
		log.Fatalf("running server failed %s", err)
	}
}
