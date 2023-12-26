package main

import (
	"fmt"
	"log"
	"time"

	"github.com/elahe-dastan/simplesurance-assignment/internal/config"
	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
	"github.com/elahe-dastan/simplesurance-assignment/internal/http"
)

const (
	banner = `
      _     _
  ___| | __| | __ _
 / _ \ |/ _| |/ _| |
|  __/ | (_| | (_| |
 \___|_|\__,_|\__,_|

`
)

func main() {
	fmt.Print(banner)

	cfg := config.New()

	// Read HitCounter data from disk or make an empty HitCounter if didn't find a file
	hc, err := hitcounter.FromFileStatic(cfg.Filename)
	if err != nil {
		log.Printf("cannot read hit counter from %s %s", cfg.Filename, err)

		hc = hitcounter.NewStatic()
	}

	// Save HitCounter data to file regularly
	tick := time.NewTicker(cfg.WriteInterval)
	defer tick.Stop()

	go func() {
		for {
			<-tick.C

			_ = hitcounter.ToFileStatic(cfg.Filename, hc)
		}
	}()

	handler := http.NewHandler(hc)

	srv := http.New()

	srv.Register(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("running server failed %s", err)
	}
}
