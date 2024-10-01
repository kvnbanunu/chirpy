package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {
    const port = 8080

    mux := http.NewServeMux()
    server := &http.Server{
        Addr: fmt.Sprintf(":%d",port),
        Handler: mux,
    }
    log.Printf("Server listening on port: %d", port)
    log.Fatal(server.ListenAndServe())
}
