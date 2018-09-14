package main

import (
  "crypto/tls"
  "net/http"
  "log"
  "time"
  "os"
)

// Following video: https://www.youtube.com/watch?v=bM6N-vgPlyQ&t=2059s
// Left off at: 6:55
// Repo for tutorial: https://github.com/dlsniper/gopherconuk

var version string = "0.0.5"

var (
  certFile = os.Getenv("CERT_FILE")
  keyFile = os.Getenv("KEY_FILE")
  serviceAddr = os.Getenv("SERVICE_ADDR")
)

func NewServer(mux *http.ServeMux, serverAddress string) *http.Server {
  // TLS setting from https://www.youtube.com/watch?v=bM6N-vgPlyQ&t=2059s
  // and https://github.com/dlsniper/gopherconuk.
  // Original source: https://blog.cloudflare.com/exposing-go-on-the-internet/
  tlsConfig := &tls.Config{
    PreferServerCipherSuites: true,
    CurvePreferences: []tls.CurveID{
      tls.CurveP256,
      tls.X25519,
    },
    MinVersion: tls.VersionTLS12,
    CipherSuites: []uint16{
      tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
      tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
      tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
      tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
      tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
      tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
    },
  }
  srv := &http.Server{
    Addr: serverAddress,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout: 120 * time.Second,
    TLSConfig: tlsConfig,
    Handler: mux,
  }
  return srv
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hello!"))})
  srv := NewServer(mux, serviceAddr)
  err := srv.ListenAndServeTLS(certFile, keyFile)
  if err != nil {
    log.Fatalf("Server failed to start: %v", err)
  }
}
