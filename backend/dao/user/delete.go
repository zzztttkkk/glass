package user

import (
	"bytes"
	"context"
	"database/sql"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/sqlx"
	"time"
)

func (Namespace) Delete(ctx context.Context, name, sec []byte) {
	var _src []byte
	var uid int64

	type Arg struct {
		Name []byte `db:"name"`
	}
	if err := op.RowColumns(ctx, "secret, id", "where name=:name and status>=0 and deleted_at=0", Arg{Name: name}, &_src, &uid); err != nil {
		if err == sql.ErrNoRows {
			return
		}
		panic(err)
	}
	if !bytes.Equal(_src, sec) {
		panic(sha.StatusError(sha.StatusBadRequest))
	}

	type Arg1 struct {
		UID int64 `db:"uid"`
	}
	op.Update(
		ctx,
		sqlx.Data{"deleted_at": time.Now().UnixNano()},
		"where id=:uid and status>=0 and deleted_at=0", Arg1{UID: uid},
	)
}
