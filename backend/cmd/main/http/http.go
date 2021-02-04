package http

import (
	"context"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/restapi"
	httplib "net/http"
	"strings"
)

func static(conf *config.Type, mux *sha.Mux) {
	isDev := conf.Env == config.EnvDevelopment
	fs := httplib.Dir(conf.Static.WebBuildPath)

	mux.NotFound = sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
		path := utils.S(ctx.Request.Path)
		if strings.HasPrefix(path, "/api") {
			ctx.Response.SetStatusCode(sha.StatusNotFound)
			return
		}
		if !isDev && strings.HasSuffix(path, ".map") {
			ctx.Response.SetStatusCode(sha.StatusNotFound)
			return
		}

		sha.ServeFileSystem(ctx, fs, path, false)
		if ctx.GetStatus() == sha.StatusNotFound && !strings.HasPrefix(path, "/static/") {
			f, _ := fs.Open("./index.html")
			stat, _ := f.Stat()
			sha.ServeFileContent(ctx, stat.Name(), stat.ModTime(), stat.Size(), f)
		}
	})
}

func Run(ctx context.Context, conf *config.Type) {
	mux := sha.NewMux(&conf.HTTP.Router, conf.GetCORSChecker())

	static(conf, mux)

	mux.AddBranch("/api", restapi.Root)
	mux.HandleDoc("get", "/api/doc")

	server := sha.New(
		ctx,
		&conf.HTTP.Server,
		sha.NewHTTP11Protocol(&conf.HTTP.HTTPProtocol),
		sha.NewWebSocketProtocol(&conf.HTTP.WebSocketProtocol),
	)
	server.Handler = mux

	mux.Print()
	server.ListenAndServe()
}
