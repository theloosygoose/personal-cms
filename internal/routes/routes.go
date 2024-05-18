package routes

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/handlers"
)

func AddRoutes(
    mux *http.ServeMux, 
    articleStore db.DB,

){
    articlehandler := handlers.ArticleHandle{ DB: articleStore }

    mux.HandleFunc("GET /api/article", articlehandler.HandleGetArticle())
}
