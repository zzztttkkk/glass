package main

import (
	"github.com/zzztttkkk/sha/utils"
	_ "glass/cmd/main/internal"
	"glass/config"
	"glass/internal"
	"glass/restapi"
	"glass/service/session"
	"strings"

	"github.com/zzztttkkk/sha"
)

func run(conf *config.Type) {
	var co func([]byte) *sha.CorsOptions
	if conf.Env == config.EnvDevelopment {
		co = func(origin []byte) *sha.CorsOptions {
			local := map[string]bool{
				"http://localhost:8089": true,
				"http://127.0.0.1:8089": true,
			}
			if local[string(origin)] {
				return (&sha.CorsOptions{AllowCredentials: true}).Init()
			}
			return nil
		}
	}

	mux := sha.NewMux(conf.HTTP.PathPrefix, co)

	mux.AddBranch("/api", restapi.Root)
	mux.HandleDoc("get", "/api/doc")

	mux.FilePath(
		conf.Static.WebBuildPath+"/static",
		"get", "/static/filename:*", false,
		sha.MiddlewareFunc(func(ctx *sha.RequestCtx, next func()) {
			fn, _ := ctx.Request.Params.Get("filename")
			if conf.Env != config.EnvDevelopment && strings.HasSuffix(utils.S(fn), ".map") {
				ctx.SetStatus(sha.StatusNotFound)
				return
			}
			next()
		}),
	)

	mux.File(
		conf.Static.WebBuildPath+"/index.html", "get", "/",
		sha.MiddlewareFunc(func(ctx *sha.RequestCtx, next func()) {
			session.New(ctx)
			next()
		}),
	)

	server := sha.Default(mux)
	if len(conf.HTTP.Host) > 0 {
		server.Host = conf.HTTP.Host
	}
	if conf.HTTP.Port > 0 {
		server.Port = conf.HTTP.Port
	}
	mux.Print()
	server.ListenAndServe()
}

func main() {
	conf := config.Type{}

	startup(&conf)

	internal.DigContainer.Provide(func() *config.Type { return &conf })
	internal.DigContainer.Invoke()

	run(&conf)
}
