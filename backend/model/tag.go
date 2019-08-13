package model

type Tag struct {
	ID    int64  `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
	CTime string `db:"ctime" json:"ctime"`
	UTime string `db:"utime" json:"utime"`
}

// リクエストする時に扱いやすい形
type RequestTags struct {
	TagIDs []int64 `json:"tag_ids"`
}
