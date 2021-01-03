package account

import (
	"github.com/zzztttkkk/sha"
	"glass/config"
	"glass/internal"
	"glass/service"
	"net/http"
)

var authCookieName string
var authHeaderName string

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			authCookieName = cfg.HTTP.AuthCookieName
			authHeaderName = cfg.HTTP.AuthHeaderName
		},
	)
}

func init() {
	type Form struct {
		Name      []byte `validator:",L=-64,F=username"`
		Password  []byte `validator:",L=6-64,R=password"`
		SetCookie bool   `validator:"cookie,optional"`
	}

	Branch.HTTPWithForm(
		"get",
		"/login",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form Form
			ctx.MustValidate(&form)
			token := service.Account.DoLogin(ctx.Context(), form.Name, form.Password)
			if len(token) < 1 {
				ctx.SetStatus(http.StatusBadRequest)
				return
			}
			if form.SetCookie {
				ctx.Response.SetCookie(authCookieName, token, &sha.CookieOptions{MaxAge: })
			}
		}),
		Form{},
	)
}
