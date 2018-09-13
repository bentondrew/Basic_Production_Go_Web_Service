package main

import (
  "net/http"
  "log"
)

// Following video: https://www.youtube.com/watch?v=bM6N-vgPlyQ&t=2059s
// Left off at: 4:55

var version string = "0.0.4"

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!"))})
  err := http.ListenAndServe(":8080", mux)
  if err != nil {
    log.Fatalf("Server failed to start: %v", err)
  }
}
