package main

import (
  "net/http"
)

var version string = "0.0.3"

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!"))})
  http.ListenAndServe(":8080", mux)
}
