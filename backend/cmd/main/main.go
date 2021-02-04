package main

import (
	"context"
	_ "embed"
	"glass/cmd/main/http"
	_ "glass/cmd/main/internal"
	"glass/config"
	"glass/internal"
	"strings"
	"sync"
)

func main() {
	var conf = config.Default()
	loadConfig(&conf)

	internal.DigContainer.Provide(func() *config.Type { return &conf })
	internal.DigContainer.Invoke()

	var wg sync.WaitGroup
	run(context.Background(), &conf, &wg)
	wg.Wait()
}

func run(ctx context.Context, conf *config.Type, wg *sync.WaitGroup) {
	if strings.Contains(*sp, "http") || strings.Contains(*sp, "all") {
		go func() {
			http.Run(ctx, conf)
			wg.Done()
		}()
		wg.Add(1)
	}
}
