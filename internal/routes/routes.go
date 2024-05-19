package routes

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/handlers"
)

func AddRoutes(
	mux *http.ServeMux,
	articleStore db.DB,
) {

	articlehandler := handlers.ArticleHandle{DB: articleStore}
	dashboardhandler := handlers.DashboardHandler{DB: articleStore}
    //API ROUTES
	mux.HandleFunc("GET /api/article", articlehandler.GetAllArticles())


    //DASHBOARD ROUTES
    mux.Handle("GET /", redirect())
	mux.HandleFunc("GET /dashboard", dashboardhandler.MainShow())
    mux.HandleFunc("GET /dashboard/writer", dashboardhandler.WriterShow())

    mux.HandleFunc("POST /dashboard/writer/new", dashboardhandler.AddArticle())
}

func redirect() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    })
}
