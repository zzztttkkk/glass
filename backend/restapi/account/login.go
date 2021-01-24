package account

import (
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/internal"
	"glass/service"
	"net/http"
	"time"
)

var authCookieName string
var authHeaderName string
var authMaxAge int64

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			authCookieName = cfg.Auth.CookieName
			authHeaderName = cfg.Auth.HeaderName
			authMaxAge = int64(cfg.Auth.TokenMaxAge)
		},
	)
}

const SixMonth = time.Hour * 24 * 30 * 6

func init() {
	type Form struct {
		Name      []byte `validator:",L=-64,F=username"`
		Password  []byte `validator:",L=6-64,R=password"`
		ByCookie  bool   `validator:"bycookie,optional"`
		KeepLogin bool   `validator:"keep,optional"`
		Referer   string `validator:"ref,optional"`
	}

	loginHandler := sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
		var form Form
		ctx.MustValidate(&form)
		token := service.Account.DoLogin(ctx, form.Name, form.Password)
		if len(token) < 1 {
			ctx.SetStatus(http.StatusBadRequest)
			return
		}

		if form.ByCookie {
			if form.KeepLogin {
				ctx.Response.SetCookie(authCookieName, token, &sha.CookieOptions{MaxAge: authMaxAge})
			} else {
				ctx.Response.SetCookie(authCookieName, token, nil)
			}
		} else {
			ctx.Response.Header.Set(authHeaderName, utils.B(token))
		}

		if len(form.Referer) > 0 {
			sha.RedirectTemporarily(form.Referer)
		}
	})

	Branch.HTTPWithForm("post", "/login", loginHandler, Form{})

	Branch.HTTPWithForm("get", "/login", loginHandler, Form{})
}
