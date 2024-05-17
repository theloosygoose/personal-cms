package server

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/middleware"
	"github.com/theloosygoose/cms-api/internal/routes"
)


func NewServer() *http.Server {
    s := http.NewServeMux() 

    routes.AddRoutes(s)

    stack := middleware.CreateStack(
        middleware.Logging,
    )

    server := http.Server{
        Addr: ":8080",
        Handler: stack(s),
    }

    return &server
}
