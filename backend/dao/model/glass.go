package model

import "github.com/zzztttkkk/sha/sqlx"

type Water struct {
	sqlx.Model
	Src      int64           `db:"src" json:"src"`
	Dst      int64           `db:"dst" json:"dst"`
	Num      int64           `db:"volume" json:"volume"`
	Category int             `db:"category" json:"category"`
	ExtID    int64           `db:"ext_id" json:"ext_id"`
	ExtInfo  sqlx.JsonObject `db:"ext_info" json:"ext_info"`
}
