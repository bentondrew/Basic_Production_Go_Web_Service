package main

import (
  "net/http"
)

var version string = "0.0.2"

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("Hello!"))})
  http.ListenAndServe(":8080", mux)
}
