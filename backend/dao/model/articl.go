package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type Article struct {
	sqlx.Model

	AuthorID int64 `db:"author_id,g=meta" json:"author_id"`
	Author   *User `db:"-" json:"author"`

	Title   sqlx.JsonBytesString `db:"title,g=meta" json:"title"`
	Summary sqlx.JsonBytesString `db:"summary,g=meta" json:"summary"`
	Content sqlx.JsonBytesString `db:"content" json:"-"`
}

func (a Article) TableName() string { return "article" }

func (a Article) TableColumns(db *x.DB) []string {
	return append(
		a.Model.TableColumns(db),
		"author_id bigint not null",
		"title varchar(256) not null",
		"summary varchar(512) default ''",
		"content text",
	)
}

type Tag struct {
	sqlx.Model
	Name        sqlx.JsonBytesString `db:"name" json:"name"`
	Color       int32                `db:"color" json:"color"`
	Description sqlx.JsonBytesString `db:"description" json:"description"`
}

func (t Tag) TableName() string { return "tag" }

func (t Tag) TableColumns(db *x.DB) []string {
	return append(
		t.Model.TableColumns(db),
		"name char(256) unique not null",
		"color int default 0",
		"description text",
	)
}

type ArticleTags struct {
	ArticleID int64 `db:"article_id"`
	TagID     int64 `db:"tag_id"`
}

func (ats ArticleTags) TableName() string { return "article_tags" }

func (ats ArticleTags) TableColumns(_ *x.DB) []string {
	return []string{
		"article_id bigint not null",
		"tag_id bigint not null",
		"primary key(article_id, tag_id)",
	}
}
