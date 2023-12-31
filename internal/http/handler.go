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

func NewHandler(hc hitcounter.HitCounter) Handler {
	return Handler{
		hc: hc,
	}
}

// ServeHTTP calls Hit with current timestamp and then returns the total number of requests in the last minute by
// calling Count and writing it to response
func (h Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.hc.Hit(time.Now().Unix())

	count := h.hc.Count()

	res.Header().Add("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)

	bytes, err := json.Marshal(count)
	if err != nil {
		log.Printf("cannot marshal count into json %s", err)

		res.WriteHeader(http.StatusInternalServerError)

		return
	}

	if _, err := res.Write(bytes); err != nil {
		log.Printf("cannot write into the response %s", err)

		res.WriteHeader(http.StatusInternalServerError)

		return
	}
}
