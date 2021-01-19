package account

import (
	"context"
	"github.com/zzztttkkk/sha/sqlx"
	"github.com/zzztttkkk/sha/utils"
	"glass/dao"
	"glass/events/account"
)

type MsgRegister struct {
	UserId int64  `json:"user_id"`
	Secret string `json:"secret"`
}

func (Namespace) DoRegister(ctx context.Context, name, password []byte) MsgRegister {
	ctx, committer := sqlx.Tx(ctx)
	defer committer()

	uid, sec := dao.User.Insert(ctx, name, password)
	account.Namespace.AfterRegister(ctx, uid)
	return MsgRegister{UserId: uid, Secret: utils.S(sec)}
}

func (Namespace) DoCheckNameExists(ctx context.Context, name []byte) bool {
	return dao.User.NameExists(ctx, name)
}
