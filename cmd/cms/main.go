package main

import (
	"log"
	"net/http"

	"github.com/theloosygoose/cms-api/internal/middleware"
	"github.com/theloosygoose/cms-api/internal/routes"
)

func main() {
    s := http.NewServeMux() 
    routes.AddRoutes(s)

    stack := middleware.CreateStack(
        middleware.Logging,
    )

    server := http.Server{
        Addr: ":8080",
        Handler: stack(s),
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Fatal("Server Failed: ", err)
    }
}
