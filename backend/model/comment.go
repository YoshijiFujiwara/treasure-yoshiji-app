package model

type Comment struct {
	ID        int64  `db:"id" json:"id"`
	UserID    int64 `db:"user_id" json:"user_id"`
	ArticleID int64  `db:"article_id" json:"article_id"`
	Body      string `db:"body" json:"body"`
}
