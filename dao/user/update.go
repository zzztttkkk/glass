package user

import (
	"context"
	"github.com/zzztttkkk/sha/sqlx"
)

func (Namespace) UpdateByID(ctx context.Context, uid int64, data sqlx.Data) {
	delete(data, "secret")
	delete(data, "id")

	type Arg struct {
		UID int64 `db:"uid"`
	}
	op.Update(ctx, data, "where id=:uid", Arg{UID: uid})
}
