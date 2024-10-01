package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/kvnbanunu/chirpy/internal/server"
)

type apiConfig struct {
    fileserverHits atomic.Int32
}

func main() {
    const port = 8080
    const filePathRoot = "."

    mux := http.NewServeMux()
    
    fsHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filePathRoot)))
    mux.Handle("/app/", fsHandler)
    mux.HandleFunc("GET /healthz", server.ReadinessHandler)

    server := &http.Server{
        Addr: fmt.Sprintf(":%d",port),
        Handler: mux,
    }

    log.Printf("Server listening on port: %d", port)
    log.Fatal(server.ListenAndServe())
}
