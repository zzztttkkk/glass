package user

import (
	"github.com/zzztttkkk/sha/sqlx"
	"glass/config"
	"glass/dao/model"
	"glass/internal"
)

var op *sqlx.Operator

func init() {
	internal.DigContainer.Append(
		func(_ *config.Type) {
			op = sqlx.NewOperator(model.User{})
			op.CreateTable(true)
		},
	)
}

type Namespace struct{}
