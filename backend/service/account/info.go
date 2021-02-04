package account

import (
	"context"
	"glass/cache"
	"glass/dao/model"
)

func (n Namespace) DoInfoByID(ctx context.Context, uid int64, dist *model.User) bool {
	return cache.User.GetByID(ctx, uid, dist)
}
