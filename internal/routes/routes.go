package routes

import (
	"net/http"

    "log"
	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/handlers"
)

func AddRoutes(s *http.ServeMux){
    conn, err := db.Connect()
    if err != nil {
        log.Panicln("Could not Connect to Database", err)
        return
    }

    d := db.NewDB(conn)

    articlehandler := &handlers.ArticleHandler{
        DB: d,
    }

    s.HandleFunc("GET /api/article", articlehandler.HandleGetArticle)
}
