package user

import (
	"context"
	"github.com/zzztttkkk/sha/groupcache"
	"glass/dao"
	"glass/dao/model"
	"strconv"
)

const _LoaderGetByID = "GBI"

type _GetByIDArg struct {
	id int64
}

func (arg _GetByIDArg) Name() string {
	return strconv.FormatInt(arg.id, 10)
}

func init() {
	NameSpace.group.Append(
		_LoaderGetByID,
		func(ctx context.Context, args groupcache.NamedArgs) (ret interface{}, err error) {
			arg, _ := args.(_GetByIDArg)
			var user model.User
			if !dao.User.GetByID(ctx, arg.id, &user) {
				return nil, nil
			}
			return &user, nil
		},
	)
}

func (ns *_Namespace) GetByID(ctx context.Context, uid int64, dist *model.User) bool {
	if err := ns.group.Do(ctx, _LoaderGetByID, dist, _GetByIDArg{id: uid}); err != nil {
		if err == groupcache.ErrEmpty {
			return false
		}
		panic(err)
	}
	return true
}

func (ns *_Namespace) DeleteByID(ctx context.Context, uid int64) {
	ns.storage.Del(ctx, ns.group.MakeKey(_LoaderGetByID, _GetByIDArg{id: uid}.Name()))
}
