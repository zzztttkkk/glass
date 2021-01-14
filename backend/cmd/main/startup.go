package main

import (
	"flag"
	"glass/config"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/zzztttkkk/sha/sqlx"
)

func startup(cfg *config.Type) {
	cp := flag.String("c", "", "config file path")
	flag.Parse()
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
				if strings.HasSuffix(info.Name(), ".json") || strings.HasSuffix(info.Name(), ".toml") {
					files = append(files, info.Name())
				}
			}
			config.FromFiles(cfg, files...)
		} else {
			config.FromFiles(cfg, *cp)
		}
	} else {
		log.Fatalln("glass: empty config")
	}

	sqlx.OpenWriteableDB(cfg.Database.DriverName, cfg.Database.WriteableURI)
	for _, uri := range cfg.Database.ReadonlyURIs {
		sqlx.OpenReadableDB(cfg.Database.DriverName, uri)
	}
}
