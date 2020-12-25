package internal

import (
	"github.com/go-sql-driver/mysql"
	"github.com/zzztttkkk/sha"
	"glass/restapi/output"
	"reflect"
)

const duplicateErr = "duplicate value"

func init() {
	sha.RecoverByType(reflect.TypeOf((*mysql.MySQLError)(nil)), func(ctx *sha.RequestCtx, v interface{}) {
		me := v.(*mysql.MySQLError)
		switch me.Number {
		case 1062: // Duplicate entry
			ctx.SetStatus(sha.StatusBadRequest)
			ctx.WriteJSON(output.Msg{Data: duplicateErr})
		default:
			ctx.SetStatus(sha.StatusInternalServerError)
		}
	})
}
