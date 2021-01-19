package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type _CommentTargetType int

const (
	CommentTargetTypeUnknown = _CommentTargetType(iota)
	CommentTargetTypeArticle
	CommentTargetTypeComment
)

type Comment struct {
	sqlx.Model

	AuthorID int64 `db:"author_id,g=meta" json:"author_id"`
	Author   *User `db:"-" json:"author"`

	ArticleID  int64              `db:"article_id" json:"article_id"`
	TargetID   int64              `db:"target_id" json:"target_id"`
	TargetType _CommentTargetType `db:"target_type" json:"target_type"`
	RootID     int64              `db:"root_id" json:"root_id"`

	Content string `db:"content" json:"content"`
}

func (c Comment) TableName() string { return "content_comment" }

func (c Comment) TableColumns(db *x.DB) []string {
	return append(
		c.Model.TableColumns(db),
		"author_id bigint not null",
		"article_id bigint not null",
		"target_id bigint not null",
		"target_type int not null",
		"root_id bigint default 0",
		"content text",
	)
}

var _ sqlx.Modeler = (*Comment)(nil)
