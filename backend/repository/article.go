package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllArticle(db *sqlx.DB) ([]model.Article, error) {
	a := make([]model.Article, 0)
	c := make([]model.Comment, 0)
	if err := db.Select(&a, `SELECT id, title, body, user_id FROM article`); err != nil {
		return nil, err
	}

	// その記事のコメントの一覧を取得する
	if err := db.Select(&c, `SELECT id, article_id, body, ctime, utime FROM article_comment`); err != nil {
		return nil, err
	}

	return a, nil
}

func FindArticle(db *sqlx.DB, id int64) (*model.Article, error) {
	a := model.Article{}
	if err := db.Get(&a, `
SELECT id, title, body, user_id FROM article WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &a, nil
}

// タグでの検索結果(単発)
func SearchByTag(db *sqlx.DB, tagId int64) (*model.Article, error) {
	a := model.Article{}
	if err := db.Get(&a, `
SELECT id, title, body, user_id, ctime, utime FROM article 
 LEFT JOIN tag_article ON article.id = tag_article.article_id 
 LEFT JOIN tags ON tag.id = tag_article.tag_id
 WHERE tag.id = ?
`, tagId); err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateArticle(db *sqlx.Tx, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article (title, body, user_id) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, a.UserID)
}

func UpdateArticle(db *sqlx.Tx, id int64, a *model.Article) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article SET title = ?, body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(a.Title, a.Body, id)
}

func DestroyArticle(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(id)
}
