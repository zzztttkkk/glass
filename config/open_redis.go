package config

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
)

var ErrRedisUnknownMode = errors.New("sha.config: unknown redis mode,[singleton,ring,cluster]")

func (t *Type) initRedisClient() {
	opts := make([]*redis.Options, 0)
	for _, url := range t.Redis.Nodes {
		option, err := redis.ParseURL(url)
		if err != nil {
			panic(err)
		}
		opts = append(opts, option)
	}

	if len(opts) < 1 {
		return
	}

	switch strings.ToLower(t.Redis.Mode) {
	case "singleton":
		t.internal.redisCli = redis.NewClient(opts[0])
		return
	case "ring":
		addrs := map[string]string{}
		var pwd string
		var name string

		for ind, opt := range opts {
			addrs[fmt.Sprintf("node.%d", ind)] = opt.Addr
			if len(opt.Password) > 0 {
				pwd = opt.Password
			}
			if len(opt.Username) > 0 {
				name = opt.Username
			}
		}

		t.internal.redisCli = redis.NewRing(&redis.RingOptions{Addrs: addrs, Password: pwd, Username: name})
		return
	case "cluster":
		var addrs []string
		var pwd string
		var name string
		for _, opt := range opts {
			addrs = append(addrs, opt.Addr)
			if len(opt.Password) > 0 {
				pwd = opt.Password
			}
			if len(opt.Username) > 0 {
				name = opt.Username
			}
		}
		t.internal.redisCli = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:    addrs,
				Username: name,
				Password: pwd,
			},
		)
		return
	default:
		panic(ErrRedisUnknownMode)
	}
}
