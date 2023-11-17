package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", MainServer)

    server_port := ":8080"

    http.ListenAndServe(server_port, nil)
}

func MainServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients- We like: Apples. ")
}
