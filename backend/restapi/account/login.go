package account

import (
	"github.com/zzztttkkk/sha"
	"glass/config"
	"glass/internal"
	"glass/restapi/output"
	"glass/service"
	"net/http"
	"time"
)

var authCookieName string

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			authCookieName = cfg.Auth.CookieName
		},
	)
}

const SixMonth = time.Hour * 24 * 30 * 6

func init() {
	type PostForm struct {
		Name     []byte `validator:",L=-64,F=username"`
		Password []byte `validator:",L=6-64,R=password"`
	}

	Branch.HTTPWithForm(
		"post",
		"/login",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form PostForm
			ctx.MustValidate(&form)
			token := service.Account.DoLogin(ctx, form.Name, form.Password)
			if len(token) < 1 {
				ctx.SetStatus(http.StatusBadRequest)
				return
			}
			output.OK(ctx, output.M{"token": token})
		}),
		PostForm{},
	)

	type GetForm struct {
		PostForm
		KeepLogin bool   `validator:"keep,optional"`
		Referer   string `validator:"ref,optional"`
	}

	Branch.HTTPWithForm(
		"get",
		"/login",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			var form GetForm
			ctx.MustValidate(&form)
			token := service.Account.DoLogin(ctx, form.Name, form.Password)
			if len(token) < 1 {
				ctx.SetStatus(http.StatusBadRequest)
				return
			}

			if !form.KeepLogin {
				ctx.Response.SetCookie(authCookieName, token, nil)
			} else {
				ctx.Response.SetCookie(authCookieName, token, &sha.CookieOptions{Expires: time.Now().Add(SixMonth)})
			}

			if len(form.Referer) > 0 {
				sha.RedirectTemporarily(form.Referer)
			}
		}),
		GetForm{},
	)
}
