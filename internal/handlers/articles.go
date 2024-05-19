package handlers

import (
	"log"
	"net/http"


	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/utils"
)

type ArticleHandle struct{ db.DB }

func (h ArticleHandle) GetAllArticles() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results, err := h.DB.DbGetArticles()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = utils.Encode(w, r, http.StatusOK, results)
		if err != nil {
			log.Println("Unable to Encode Reponse to JSON")
			return
		}
	})
}

func (h ArticleHandle) GetOneArticle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Get Single Article
	})
}

