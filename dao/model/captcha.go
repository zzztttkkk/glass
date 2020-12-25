package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type CaptchaTxt struct {
	sqlx.Model
	Txt string `db:"txt"`
}

func (CaptchaTxt) TableName() string { return "captcha_txt" }

func (c CaptchaTxt) TableColumns(db *x.DB) []string {
	return append(
		c.Model.TableColumns(db),
		"txt varchar(64) not null",
	)
}
