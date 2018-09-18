package home

import (
  "net/http"
  "log"
)


type Handlers struct {
  logger *log.Logger
}


func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
  h.logger.Println("Home request processed.")
  w.Header().Set("Content-Type", "text/plain; charset=utf-8")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello!"))
}


func NewHandlers(logger *log.Logger) *Handlers {
  return &Handlers{
    logger: logger,
  }
}
