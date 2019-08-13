package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/httputil"
	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/service"
	"log"
	"net/http"
	"strconv"
)

type Comment struct {
	db *sqlx.DB
}

func NewComment(db *sqlx.DB) *Comment {
	return &Comment{db: db}
}

func (c *Comment) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	articleId, ok := vars["id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	aid, err := strconv.ParseInt(articleId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	comments, err := repository.AllCommentOfTheArticle(c.db, aid)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, comments, nil
}

func (c *Comment) Show(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	commentId, ok := vars["comment_id"]
	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}

	cid, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	comment, err := repository.FindComment(c.db, cid)
	if err != nil && err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, comment, nil
}

func (c *Comment) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	// クエリパラメータの取得
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	aid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	newComment := &model.Comment{}
	newComment.ArticleID = aid

	if err := json.NewDecoder(r.Body).Decode(&newComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	user, err := httputil.GetUserFromContext(r.Context())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("userid")
	fmt.Println(user.ID)
	newComment.UserID = user.ID

	commentService := service.NewCommentService(c.db)
	newId, err := commentService.Create(newComment)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newComment.ID = newId

	return http.StatusCreated, newComment, nil
}


func (c *Comment) Update(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	commentId, ok := vars["comment_id"]

	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	cid, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	reqComment := &model.Comment{}
	if err := json.NewDecoder(r.Body).Decode(&reqComment); err != nil {
		return http.StatusBadRequest, nil, err
	}

	commentService := service.NewCommentService(c.db)
	err = commentService.Update(cid, reqComment)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}

func (c *Comment) Destroy(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	vars := mux.Vars(r)
	commentId, ok := vars["comment_id"]

	if !ok {
		return http.StatusBadRequest, nil, &httputil.HTTPError{Message: "invalid path parameter"}
	}
	cid, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	commentService := service.NewCommentService(c.db)
	err = commentService.Destroy(cid)
	if err != nil && errors.Cause(err) == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	} else if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusNoContent, nil, nil
}