package main

import (
  "net/http"
  "log"
  "os"
  "server"
)

// Following video: https://www.youtube.com/watch?v=bM6N-vgPlyQ&t=2059s
// Left off at: 6:55
// Repo for tutorial: https://github.com/dlsniper/gopherconuk

var version string = "0.0.6"

var (
  certFile = os.Getenv("CERT_FILE")
  keyFile = os.Getenv("KEY_FILE")
  serviceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!"))})
  srv := server.Nwe(mux, serviceAddr)
  err := srv.ListenAndServeTLS(certFile, keyFile)
  if err != nil {
    log.Fatalf("Server failed to start: %v", err)
  }
}
