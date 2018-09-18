package main

import (
  "net/http"
  "log"
  "os"
  "service/server"
  "service/home"
)

// Following video: https://www.youtube.com/watch?v=bM6N-vgPlyQ&t=2059s
// Left off at: 13.24
// Repo for tutorial: https://github.com/dlsniper/gopherconuk
// Test building from http://blog.setale.me/2018/01/06/Go-Tests-multistage-docker-builds-pipeline/


var version string = "0.0.8"


var (
  certFile = os.Getenv("CERT_FILE")
  keyFile = os.Getenv("KEY_FILE")
  serviceAddr = os.Getenv("SERVICE_ADDR")
)


func main() {
  logger := log.New(os.Stdout, "Basic Web Service ", log.LstdFlags|log.Lshortfile)
  h := home.NewHandlers(logger)
  mux := http.NewServeMux()
  h.SetupRoutes(mux)
  srv := server.New(mux, serviceAddr)
  logger.Println("Server starting.")
  err := srv.ListenAndServeTLS(certFile, keyFile)
  if err != nil {
    logger.Fatalf("Server failed to start: %v", err)
  }
}
