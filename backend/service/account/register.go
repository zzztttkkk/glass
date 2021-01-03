package account

import (
	"context"
	"github.com/zzztttkkk/sha/sqlx"
	"glass/dao"
	"glass/events"
)

type MsgRegister struct {
	UserId int64                `json:"user_id"`
	Secret sqlx.JsonBytesString `json:"secret"`
}

func (Namespace) DoRegister(ctx context.Context, name, password []byte) MsgRegister {
	ctx, committer := sqlx.Tx(ctx)
	defer committer()

	uid, sec := dao.User.Insert(ctx, name, password)
	events.Account.AfterRegister(ctx, uid)
	return MsgRegister{UserId: uid, Secret: sec}
}

func (Namespace) DoCheckNameExists(ctx context.Context, name []byte) bool {
	return dao.User.NameExists(ctx, name)
}
