package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/metrics", metricsHandler)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func metricsHandler(writer http.ResponseWriter, request *http.Request) {
    writer.WriteHeader(http.StatusOK)
    fmt.Fprint(writer, "Metrics endpoint reached successfully.")
}