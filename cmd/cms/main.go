package main

import (
	"log"
    _"net/http"

    "github.com/joho/godotenv"

	"github.com/theloosygoose/cms-api/internal/server"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Panic("Error Loading .env file")
    }

    server := server.NewServer()
    log.Println("Making New Server")

    err = server.ListenAndServe()
    if err != nil {
        log.Fatal("Server Failed: ", err)
    } else {
        log.Println("Server Starting")
    }
}
