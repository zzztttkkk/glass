package main

import (
	"github.com/zzztttkkk/sha"
	_ "glass/cmd/main/internal"
	"glass/config"
	"glass/internal"
	"glass/restapi"
)

func loadHttpServicesAndRun(conf *config.Type) {
	opt := &conf.HTTP.CorsOptions
	if opt.CheckOrigin == nil {
		opt = nil
	}
	mux := sha.NewMux(conf.HTTP.PathPrefix, opt)

	mux.AddBranch("/api", restapi.Root)

	server := sha.Default(mux)
	if len(conf.HTTP.Host) > 0 {
		server.Host = conf.HTTP.Host
	}
	if conf.HTTP.Port > 0 {
		server.Port = conf.HTTP.Port
	}
	mux.Print(false, false)
	server.ListenAndServe()
}

func main() {
	conf := config.Type{}

	loadConfigAndConnectDatabase(&conf)

	internal.DigContainer.Provide(func() *config.Type { return &conf })
	internal.DigContainer.Invoke()

	loadHttpServicesAndRun(&conf)
}
