package account

import (
	"context"
	"github.com/zzztttkkk/sha/auth"
	"github.com/zzztttkkk/sha/sqlx"
	"github.com/zzztttkkk/sha/utils"
	"glass/dao"
	"glass/dao/model"
	"glass/events/account"
)

func (Namespace) DoUnregister(ctx context.Context, secret []byte) {
	ctx, committer := sqlx.Tx(ctx)
	defer committer()

	user := auth.MustAuth(ctx).(*model.User)

	account.Namespace.BeforeUnregister(ctx, user)
	dao.User.Delete(ctx, utils.B(user.Name), secret)
}
