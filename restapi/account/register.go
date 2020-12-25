package account

import (
	"github.com/zzztttkkk/sha"
	"glass/restapi/output"
	"glass/service"
)

func init() {
	type NameForm struct {
		Name []byte `validator:",L=-64,F=username"`
	}

	Branch.HTTPWithForm(
		"get",
		"/name_exists",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form NameForm
			ctx.MustValidate(&form)
			output.OK(ctx, service.Account.DoCheckNameExists(ctx.Context(), form.Name))
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
			output.OK(ctx, service.Account.DoRegister(ctx.Context(), form.Name, form.Password))
		}),
		Form{},
	)

}
