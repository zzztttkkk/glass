package model

import (
	x "github.com/jmoiron/sqlx"
	"github.com/zzztttkkk/sha/sqlx"
)

type Water struct {
	sqlx.Model
}

var _ sqlx.Modeler = (*Water)(nil)

func (w Water) TableName() string { return "account_water" }

func (w Water) TableColumns(db *x.DB) []string {
	return append(
		w.Model.TableColumns(db),
		"",
	)
}
