package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type Tag struct {
	sqlx.Model
	Name        string `db:"name" json:"name"`
	Color       int32  `db:"color" json:"color"`
	Description string `db:"description" json:"description"`
}

var _ sqlx.Modeler = (*Tag)(nil)

func (t Tag) TableName() string { return "content_tag" }

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

var _ sqlx.Modeler = (*ArticleTags)(nil)

func (ats ArticleTags) TableName() string { return "content_article_tags" }

func (ats ArticleTags) TableColumns(_ *x.DB) []string {
	return []string{
		"article_id bigint not null",
		"tag_id bigint not null",
		"primary key(article_id, tag_id)",
	}
}
