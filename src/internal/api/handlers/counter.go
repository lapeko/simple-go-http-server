package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type counterHandler struct {
	counter atomic.Uint64
}

func NewCounter() http.Handler {
	return &counterHandler{}
}

// GetCounter godoc
// @Summary Get counter
// @Tags counter
// @Produce plain
// @Success 200 {integer} integer
// @Router /counter [get]
func swaggerGetCounter() {}

// IncrementCounter godoc
// @Summary Increment counter
// @Tags counter
// @Success 204 "No Content"
// @Router /counter [post]
func swaggerIncrementCounter() {}

func (c *counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%d", c.counter.Load())
		}
	case http.MethodPost:
		{
			c.counter.Add(1)
			w.WriteHeader(http.StatusNoContent)
		}
	default:
		{
			w.Header().Set("Allow", "GET, POST")
			w.Write([]byte("Method not allowed"))
		}
	}
}
