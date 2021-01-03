package account

import (
	"context"
	"github.com/zzztttkkk/sha/auth"
	"github.com/zzztttkkk/sha/sqlx"
	"glass/dao"
	"glass/dao/model"
	"glass/events"
)

func (Namespace) DoUnregister(ctx context.Context, secret []byte) {
	ctx, committer := sqlx.Tx(ctx)
	defer committer()

	user := auth.MustAuth(ctx).(*model.User)

	events.Account.BeforeUnregister(ctx, user)
	dao.User.Delete(ctx, user.Name, secret)
}
