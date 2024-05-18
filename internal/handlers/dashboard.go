package handlers

import (
	"net/http"

	"github.com/theloosygoose/cms-api/internal/db"
	"github.com/theloosygoose/cms-api/internal/utils"
	"github.com/theloosygoose/cms-api/internal/views/dashboard"
)




type DashboardHandler struct{ db.DB }


func (h DashboardHandler) DashboardMainShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        utils.TemplRender(w, r, dashboard.Home())
    })
}

func (h DashboardHandler) DashboardWriterShow() http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        utils.TemplRender(w, r, dashboard.Writer())
    })
}
