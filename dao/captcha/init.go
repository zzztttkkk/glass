package captcha

import (
	"fmt"
	"github.com/zzztttkkk/sha/sqlx"
	"glass/config"
	"glass/dao/model"
	"glass/internal"
)

var op *sqlx.Operator

var insertSql string

func init() {
	internal.DigContainer.Append(
		func(_ *config.Type) {
			op = sqlx.NewOperator(model.CaptchaTxt{})
			op.CreateTable(true)
			insertSql = fmt.Sprintf("insert into %s (created_at,txt) values(?,?)", op.TableName())
		},
	)
}

type Namespace struct{}
