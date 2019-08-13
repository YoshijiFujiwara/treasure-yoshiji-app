package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/voyagegroup/treasure-app/model"
)

func AllComment(db *sqlx.DB) ([]model.Comment, error) {
	c := make([]model.Comment, 0)

	if err := db.Select(&c, `SELECT id, user_id, article_id, body, ctime, utime FROM article_comment`); err != nil {
		return nil, err
	}
	return c, nil
}

func AllCommentOfTheArticle(db *sqlx.DB, articleId int64) ([]model.Comment, error) {
	c := make([]model.Comment, 0)

	if err := db.Select(&c, `SELECT id, user_id, article_id, body, ctime, utime FROM article_comment WHERE article_id = ?`, articleId); err != nil {
		return nil, err
	}
	return c, nil
}

func FindComment(db *sqlx.DB, id int64) (*model.Comment, error) {
	c := model.Comment{}
	if err := db.Get(&c, `
SELECT id, user_id, article_id, body, ctime, utime FROM article_comment WHERE id = ?
`, id); err != nil {
		return nil, err
	}
	return &c, nil
}


func CreateComment(db *sqlx.Tx, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`
INSERT INTO article_comment (user_id, article_id, body) VALUES (?, ?, ?)
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.UserID, c.ArticleID, c.Body)
}

func UpdateComment(db *sqlx.Tx, id int64, c *model.Comment) (sql.Result, error) {
	stmt, err := db.Prepare(`
UPDATE article_comment SET body = ? WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(c.Body, c.ID)
}

func DestroyComment(db *sqlx.Tx, id int64) (sql.Result, error) {
	stmt, err := db.Prepare(`
DELETE FROM article_comment WHERE id = ?
`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	fmt.Println("delete id")
	fmt.Println(id)

	return stmt.Exec(id)
}


