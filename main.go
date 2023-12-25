package main

import (
	"fmt"
	"log"

	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
	"github.com/elahe-dastan/simplesurance-assignment/internal/http"
)

const banner = `
      _     _
  ___| | __| | __ _
 / _ \ |/ _| |/ _| |
|  __/ | (_| | (_| |
 \___|_|\__,_|\__,_|

`

func main() {
	fmt.Print(banner)

	hc := hitcounter.NewStatic()

	handler := http.NewHandler(hc)

	srv := http.New()

	srv.Register(handler)

	if err := srv.Run(); err != nil {
		log.Fatalf("running server failed %s", err)
	}
}
