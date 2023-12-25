package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/elahe-dastan/simplesurance-assignment/internal/hitcounter"
)

type Handler struct {
	hc hitcounter.HitCounter
}

func (h Handler) Handle(req *http.Request, res http.ResponseWriter) {
	h.hc.Hit(time.Now().Unix())

	count := h.hc.Count()

	bytes, err := json.Marshal(count)
	if err != nil {
		log.Printf("cannot marshal count into json %s", err)

		return
	}

	if _, err := res.Write(bytes); err != nil {
		log.Printf("cannot write into the response %s", err)

		return
	}

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)
}
