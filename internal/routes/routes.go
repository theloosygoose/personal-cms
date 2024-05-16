package routes

import (
    "net/http"
)

func AddRoutes(s *http.ServeMux) {
    s.HandleFunc("GET /api/article")
}
