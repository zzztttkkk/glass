package middleware

import (
	"github.com/zzztttkkk/sha"
	"glass/service"
)

var RequireLogin = sha.MiddlewareFunc(func(ctx *sha.RequestCtx, next func()) {
	_, err := service.Account.Auth(ctx)
	if err != nil {
		panic(err)
	}
	next()
})
