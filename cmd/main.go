package main

import (
    "net/http"
    "log"
)

func main() {
    s := http.NewServeMux() 

    server := http.Server{
        Addr: ":8080",
        Handler: s,
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Server Failed: ", err)
    }
}
