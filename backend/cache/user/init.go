package user

import (
	"github.com/zzztttkkk/sha/groupcache"
	"glass/config"
	"glass/internal"
)

type _Namespace struct {
	group   *groupcache.Group
	storage groupcache.Storage
}

var NameSpace = &_Namespace{group: groupcache.Simple()}

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			NameSpace.storage = cfg.GetCacheStorage()

			NameSpace.group.
				Init("user", cfg.Cache.Expires, cfg.Cache.MaxWaitWhenCaching).
				SetStorage(NameSpace.storage).
				SetStoragePrefix("cache")
		},
	)
}
