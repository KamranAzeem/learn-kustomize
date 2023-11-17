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
    // keep changing the name of city to mimic change in the code.
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients! We are located in: Oslo")
}
