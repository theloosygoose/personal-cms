package handlers

import (
	"net/http"
    "log"
    "io"

	"github.com/gomarkdown/markdown"
	"github.com/k3a/html2text"

	"github.com/theloosygoose/cms-api/internal/types"
	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/utils"
	"github.com/theloosygoose/cms-api/internal/views/dashboard"
)




type DashboardHandler struct{ db.DB }


func (h DashboardHandler) MainShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        utils.TemplRender(w, r, dashboard.Home())
    })
}

func (h DashboardHandler) WriterShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        utils.TemplRender(w, r, dashboard.Writer())
    })
}

func (h DashboardHandler) AddArticle() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        //Add an Article
        r.ParseMultipartForm(10)

		file, _, err := r.FormFile("article-file")
		if err != nil {
			log.Println(err)
			return
		}
        mdBytes, err := io.ReadAll(file)

        html := string(markdown.ToHTML(mdBytes, nil, nil))
        plain_text := html2text.HTML2Text(html)

        details := types.Article{
            Title: r.FormValue("article-title"),
            Body: html,
            PlainText: plain_text,
        }

        err = h.DB.DbAddArticle(details)
        if err != nil {
            log.Println("Error adding article to Database", err)
        }

        http.Redirect(w, r, "/dashboard/writer", http.StatusCreated)
    })
}
