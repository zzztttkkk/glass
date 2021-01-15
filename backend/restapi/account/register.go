package account

import (
	"glass/restapi/output"
	"glass/service"

	"github.com/zzztttkkk/sha"
)

func init() {
	type NameForm struct {
		Name []byte `validator:",L=-64,F=username"`
	}

	Branch.HTTPWithForm(
		"get",
		"/exists",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form NameForm
			ctx.MustValidate(&form)
			output.OK(ctx, service.Account.DoCheckNameExists(ctx, form.Name))
		}),
		NameForm{},
	)

	type Form struct {
		NameForm
		Password []byte `validator:",L=6-64,R=password"`
	}

	Branch.HTTPWithForm(
		"post",
		"/register",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form Form
			ctx.MustValidate(&form)
			output.OK(ctx, service.Account.DoRegister(ctx, form.Name, form.Password))
		}),
		Form{},
	)
}
