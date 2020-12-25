package account

import "github.com/zzztttkkk/sha"

func init() {
	type Form struct {
		Name     []byte `validator:",L=-64,F=username"`
		Password []byte `validator:",L=6-64,R=password"`
	}

	Branch.HTTPWithForm(
		"get",
		"/login",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {

		}),
		Form{},
	)
}
