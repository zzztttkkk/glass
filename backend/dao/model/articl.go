package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type Article struct {
	sqlx.Model
	HID string `db:"hid" json:"hid"`

	AuthorID int64 `db:"author_id,g=meta" json:"author_id"`
	Author   *User `db:"-" json:"author"`

	Title   string `db:"title,g=meta" json:"title"`
	Summary string `db:"summary,g=meta" json:"summary"`
	Content string `db:"content" json:"-"`
}

var _ sqlx.Modeler = (*Article)(nil)

func (a Article) TableName() string { return "content_article" }

func (a Article) TableColumns(db *x.DB) []string {
	return append(
		a.Model.TableColumns(db),
		"hid char(64) unique not null",
		"author_id bigint not null",
		"title varchar(256) not null",
		"summary varchar(512) default ''",
		"content text",
	)
}
