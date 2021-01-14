package main

import (
	_ "glass/cmd/main/internal"
	"glass/config"
	"glass/internal"
	"glass/restapi"

	"github.com/zzztttkkk/sha"
)

func run(conf *config.Type) {
	opt := &conf.HTTP.CorsOptions
	if opt.CheckOrigin == nil {
		opt = nil
	}

	mux := sha.NewMux(conf.HTTP.PathPrefix, opt)

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

func main() {
	conf := config.Type{}

	startup(&conf)

	internal.DigContainer.Provide(func() *config.Type { return &conf })
	internal.DigContainer.Invoke()

	run(&conf)
}
