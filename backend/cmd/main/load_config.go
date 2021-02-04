package main

import (
	"flag"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/zzztttkkk/sha/sqlx"
)

var cp = flag.String("c", "", "config file path")
var sp = flag.String("s", "all", "services to run, such as `http, all`")

func init() {
	flag.Parse()
}

func isConfFile(fn string) bool {
	return strings.HasPrefix(fn, "glass.") && (strings.HasSuffix(fn, ".json") || strings.HasSuffix(fn, ".toml"))
}

func loadConfig(cfg *config.Type) {
	if len(*cp) > 0 {
		stat, err := os.Stat(*cp)
		if err != nil {
			panic(err)
		}

		if stat.IsDir() {
			infos, err := ioutil.ReadDir(*cp)
			if err != nil {
				panic(err)
			}
			var files []string
			for _, info := range infos {
				if info.IsDir() {
					continue
				}
				if isConfFile(info.Name()) {
					files = append(files, *cp+"/"+info.Name())
				}
			}
			utils.Conf.LoadFromFiles(cfg, files...)
		} else {
			if !isConfFile(*cp) {
				log.Fatalf("glass.loadConfig: `%s` is not a `json` or `toml` file\n", *cp)
			}
			utils.Conf.LoadFromFiles(cfg, *cp)
		}
	} else {
		log.Fatalln("glass.loadConfig: empty config")
	}

	sqlx.OpenWriteableDB(cfg.Database.DriverName, cfg.Database.WriteableURI)
	for _, uri := range cfg.Database.ReadonlyURIs {
		sqlx.OpenReadableDB(cfg.Database.DriverName, uri)
	}
}
