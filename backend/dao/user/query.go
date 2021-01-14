package user

import (
	"context"
	"database/sql"
	"glass/dao/model"
)

func (Namespace) NameExists(ctx context.Context, name []byte) bool {
	type Arg struct {
		Name []byte `db:"name"`
	}
	var uid int64
	if err := op.RowColumns(ctx, "id", "where name=:name", Arg{Name: name}, &uid); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}
	return true
}

func (Namespace) Auth(ctx context.Context, name, pwd []byte) (uid int64) {
	type Arg struct {
		Name []byte `db:"name"`
	}
	var _pwd []byte
	if err := op.RowColumns(
		ctx,
		"id,password",
		"where name=:name and deleted_at=0 and status>=0", Arg{Name: name},
		&uid, &_pwd,
	); err != nil {
		if err == sql.ErrNoRows {
			return -1
		}
		panic(err)
	}
	if pwdHashPool.Equal(pwd, _pwd) {
		return uid
	}
	return -1
}

func (Namespace) InfoByID(ctx context.Context, uid int64) *model.User {
	var user model.User
	type Arg struct {
		UID int64 `db:"uid"`
	}
	if err := op.FetchOne(
		ctx, "info",
		"where id=:uid and deleted_at=0 and status>=0", Arg{UID: uid},
		&user,
	); err != nil {
		return nil
	}
	return &user
}
