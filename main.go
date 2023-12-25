package main

import (
	"fmt"
	"log"
	"time"

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
	fn = "state.json"
)

func main() {
	fmt.Print(banner)

	hc, err := hitcounter.FromFileStatic(fn)
	if err != nil {
		log.Printf("cannot read hit counter from %s %s", fn, err)

		hc = hitcounter.NewStatic()
	}

	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()

	go func() {
		for {
			<-tick.C

			_ = hitcounter.ToFileStatic(fn, hc)
		}
	}()

	handler := http.NewHandler(hc)

	srv := http.New()

	srv.Register(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("running server failed %s", err)
	}
}
