package main

import (
    "io"
    "net/http"
    "os"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func main() {
    healthPort := "8000"
    port := os.Getenv("HEALTH_PORT")
    if port != "" {
        healthPort = os.Getenv("HEALTH_PORT")
    }
    http.HandleFunc("/health", hello)
    http.ListenAndServe(":" + healthPort, nil)
}

