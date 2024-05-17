package main

import (
	"log"
    _"net/http"

	"github.com/theloosygoose/cms-api/internal/server"
)

func main() {
    server := server.NewServer()

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Server Failed: ", err)
    }
}
