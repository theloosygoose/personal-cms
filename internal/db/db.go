package db

import (
	"log"

	"github.com/theloosygoose/cms-api/internal/types"
)

func (d *DB) DbGetArticles() ([]types.Article, error) {
    var articles []types.Article

    query := `SELECT * FROM articles;`

    results, err := d.Query(query)
    if err != nil {
        log.Println("Error Running Query to Get All Articles", err)
        return nil, err
    }

    for results.Next() {
        var article types.Article

        err = results.Scan(&article.Id, &article.Title, &article.Body, &article.Created_at)
        if err != nil {
            log.Println("Error Scanning Results")
            return nil, err
        }

        articles = append(articles, article)
    }

    return articles, nil
}

func (d *DB) DbAddArticle(a types.Article) error {
    query := `INSERT INTO articles
    (title, body) VALUES (?,?);`

    result, err := d.Exec(query, &a.Title, &a.Body)

    if err != nil {
        log.Println("Error Adding to Database")
        return err
    }
    log.Println("Added to Database at:")
    log.Println(result.RowsAffected())

    return nil
}
