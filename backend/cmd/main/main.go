package main

import (
	_ "embed"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/utils"
	_ "glass/cmd/main/internal"
	"glass/config"
	"glass/dist"
	"glass/internal"
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

	const faviconFilename = "favicon.ico"
	favicon := httplib.FS(dist.Favicon)
	mux.HTTP("get", "/favicon.ico", sha.RequestHandlerFunc(func(ctx *sha.RequestCtx) {
		f, _ := favicon.Open(faviconFilename)
		stat, _ := f.Stat()
		sha.ServeFileContent(ctx, faviconFilename, stat.ModTime(), stat.Size(), f)
	}))
}

func http(conf *config.Type) {
	mux := sha.NewMux(conf.HTTP.PathPrefix, conf.GetCORSChecker())

	static(conf, mux)

	mux.AddBranch("/api", restapi.Root)
	mux.HandleDoc("get", "/api/doc")

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

func run(conf *config.Type) {
	http(conf)
}

func main() {
	conf := config.Type{}

	startup(&conf)

	internal.DigContainer.Provide(func() *config.Type { return &conf })
	internal.DigContainer.Invoke()

	run(&conf)
}
