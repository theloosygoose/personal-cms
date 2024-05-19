package handlers

import (
	"log"
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/utils"
	"github.com/theloosygoose/cms-api/internal/views/responses"
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
        id := r.PathValue("id")
        contype := r.Header.Get("Content-Type")

        res, err := h.DbGetSingleArticle(id)
        if err != nil {
            log.Println("Could Not Get Single Article")
            w.WriteHeader(http.StatusInternalServerError)
        }
        var f int
        if res.PlainText == ""{
            f = http.StatusNoContent
        }
        f = http.StatusOK

        if contype == "application/json" {
            utils.Encode(w, r, f, res)

        } else if contype =="text/html; charset=utf-8"{
            utils.TemplRender(w, r, responses.ArticleShow(res))
        } else {
            log.Println("Content-Type not Accepted")
        }

	})
}

