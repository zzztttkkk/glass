package model

import (
	"context"
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/auth"
	"github.com/zzztttkkk/sha/sqlx"
)

type User struct {
	sqlx.Model

	// info
	Name   string `db:"name,g=info|pub_info" json:"name"`
	Alias  string `db:"alias,g=info|pub_info"`
	Avatar string `db:"avatar,g=info|pub_info" json:"avatar"`

	// secret
	Password string `db:"password" json:"-"`
	Secret   string `db:"secret" json:"-"`
}

var _ auth.Subject = (*User)(nil)
var _ sqlx.Modeler = (*User)(nil)

func (u *User) GetID() int64 { return u.ID }

func (u *User) Info(ctx context.Context) interface{} { return nil }

func (u User) TableName() string { return "account_user" }

func (u User) TableColumns(db *x.DB) []string {
	return append(
		u.Model.TableColumns(db),
		"name char(64) unique not null",
		"alias varchar(64) default ''",
		"avatar varchar(256) default ''",
		"password char(64) not null",
		"secret text not null",
	)
}
