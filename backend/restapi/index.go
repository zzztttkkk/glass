package restapi

import (
	"github.com/zzztttkkk/sha"
	"glass/restapi/account"
	"glass/service"
	"glass/service/session"
	"image/png"
)

var Root = sha.NewBranch()

func init() {
	Root.Use(
		sha.MiddlewareFunc(func(ctx *sha.RequestCtx, next func()) {
			session.New(ctx)
			next()
		}),
	)

	Root.AddBranch("/account", account.Branch)

	Root.HTTP(
		"get", "/captcha.png",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			img := session.New(ctx).CaptchaGenPNG(ctx)
			ctx.Response.Header.SetContentType(sha.MIMEPng)
			if err := png.Encode(ctx, img); err != nil {
				panic(err)
			}
		}),
	)

	Root.HTTP(
		"get", "/built-time.txt",
		sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
			ctx.Response.Header.SetContentType(sha.MIMEText)
			_, _ = ctx.WriteString(service.BuiltTime(nil))
		}),
	)
}
