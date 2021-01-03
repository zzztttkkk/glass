package session

import (
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/captcha"
	"image/png"
)

func (s Type) CaptchaGen(ctx *sha.RequestCtx) {
	ctx.Response.Header.SetContentType(sha.MIMEPng)
	img := captcha.RenderOneFont("*", "", nil)
	if err := png.Encode(ctx, img); err != nil {
		panic(err)
	}
}
