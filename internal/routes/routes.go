package routes

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/handlers"
)

func AddRoutes(s *http.ServeMux){

    articlehandler := handlers.ArticleHandle{}

    s.HandleFunc("GET /api/article", articlehandler.HandleGetArticle)
}
