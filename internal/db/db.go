package db

import (
	"log"

	"github.com/theloosygoose/cms-api/internal/types"

)

func (d DB) DbGetArticles() ([]types.Article, error) {
    var articles []types.Article

    query := `SELECT id, title, body, plaintext, created_on FROM articles;`

    results, err := d.Query(query)
    if err != nil {
        log.Println("Error Running Query to Get All Articles", err)
        return nil, err
    }

    for results.Next() {
        var article types.Article

        err = results.Scan(&article.Id, &article.Title, &article.Body, &article.PlainText, &article.Created_at)
        if err != nil {
            log.Println("Error Scanning Results")
            return nil, err
        }

        articles = append(articles, article)
    }

    return articles, nil
}

func (d DB) DbGetSingleArticle(id string) (types.Article, error) {
    query := `SELECT id, title, body, plaintext, created_on 
    FROM articles WHERE id=?;`

    var a types.Article

    err := d.QueryRow(query, id).Scan(&a.Id, &a.Title, &a.Body, &a.PlainText, &a.Created_at)
    if err != nil {
        log.Println("Error Getting Scanning Article", err)
        return a, err
    }

    return a, nil 
}

func (d DB) DbAddArticle(a types.Article) error {
    query := `INSERT INTO articles (title, body, plaintext) VALUES (?,?,?);`

    result, err := d.Exec(query, &a.Title, &a.Body, &a.PlainText)
    if err != nil {
        log.Println("Error Adding to Database")
        return err
    }

    log.Println("Added to Database at:")
    log.Println(result.RowsAffected())

    return nil
}

