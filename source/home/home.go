package home

import (
  "net/http"
  "log"
  "time"
)


type Handlers struct {
  logger *log.Logger
}


func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello!"))
}


func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    startTime := time.Now()
    next(w, r)
    h.logger.Printf("Request processed in %s seconds.\n", time.Now().Sub(startTime))
  }
}


func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
  mux.HandleFunc("/", h.Logger(h.Home)
}


func NewHandlers(logger *log.Logger) *Handlers {
  return &Handlers{
    logger: logger,
  }
}
