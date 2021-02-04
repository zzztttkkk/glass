package account

import (
	"github.com/zzztttkkk/sha"
	"glass/dao/model"
	"glass/restapi/output"
	"glass/service"
)

func init() {
	type Form struct {
		ID int64 `validator:"id,v=1-"`
	}

	Branch.HTTPWithForm(
		"get", "/info",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form Form
			ctx.MustValidate(&form)

			var user model.User
			if !service.Account.DoInfoByID(ctx, form.ID, &user) {
				panic(sha.StatusError(sha.StatusNotFound))
			}
			output.OK(ctx, &user)
		}),
		Form{},
	)
}
