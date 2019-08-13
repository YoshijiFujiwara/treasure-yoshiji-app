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
