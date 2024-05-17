package handlers

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
)

type ArticleHandler struct {
    db db.DB
}

func (h *ArticleHandler) HandleGetArticle(w http.ResponseWriter, r *http.Request){
    results, err := h.db.DbGetArticles()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    }

    w.WriteHeader(http.StatusOK)

    w.Write([]byte(results))
}
