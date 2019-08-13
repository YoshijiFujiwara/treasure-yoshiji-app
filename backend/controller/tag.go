package controller

import (
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Tag struct {
	db *sqlx.DB
}

func NewTag(db *sqlx.DB) *Article {
	return &Article{db: db}
}

func (t *Tag) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	// todo tagの作成

	return 400, nil, nil
}
