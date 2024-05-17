package handlers

import (
	"encoding/json"
	"net/http"
    "log"

	"github.com/theloosygoose/cms-api/internal/db"
)

type ArticleHandler struct {
    DB db.DB
}

func (h *ArticleHandler) HandleGetArticle(w http.ResponseWriter, r *http.Request){
    results, err := h.DB.DbGetArticles()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(results)

    if err != nil {
        log.Println("Unable to Encode Reponse to JSON")
        return
    }
}
