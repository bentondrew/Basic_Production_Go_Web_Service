package main

import (
  "fmt"
  "net/http"
)

var version string = "0.0.2"

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){w.Write("Hello!")})
  http.ListenAndServe(":8080", mux)
}
