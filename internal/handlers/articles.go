package handlers

import (
	"encoding/json"
	"net/http"
    "log"

	"github.com/theloosygoose/cms-api/internal/db"
)

type ArticleHandle struct {}

func (h ArticleHandle) HandleGetArticle(w http.ResponseWriter, r *http.Request){
    conn, err := db.Connect()
    if err != nil {
        log.Println("Could not Connect to DB")
    }
    d := db.NewDB(conn)

    err = d.Ping()
    if err != nil {
        log.Panicln("DB Ping Error")
        return
    }

    results, err := d.DbGetArticles()
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

    db.CloseConnection(d)
}
