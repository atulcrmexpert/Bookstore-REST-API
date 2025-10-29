package main

import (
    "fmt"
    "log"
    "net/http"
    "example.com/bookstore/config"
    "example.com/bookstore/routes"
)

func main() {
    cfg := config.Load()
    mux := routes.SetupRoutes()

    addr := fmt.Sprintf(":%d", cfg.Port)
    log.Printf("Starting Bookstore API on %s\n", addr)
    if err := http.ListenAndServe(addr, mux); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}
