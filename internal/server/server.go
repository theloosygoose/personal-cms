package server

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/middleware"
	"github.com/theloosygoose/cms-api/internal/routes"
)


func NewServer(database db.DB) *http.Server {
    s := http.NewServeMux() 

    routes.AddRoutes(s, database)

    stack := middleware.CreateStack(
        middleware.Logging,
    )

    server := http.Server{
        Addr: ":8080",
        Handler: stack(s),
    }

    return &server
}
