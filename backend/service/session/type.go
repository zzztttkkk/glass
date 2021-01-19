package session

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/rs/xid"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/internal"
	"strconv"
	"time"
)

var cli redis.Cmdable
var cookieName string
var headerName string
var headerExpireName string
var maxAge time.Duration
var maxAgeSeconds int64
var qOfMaxAge int64
var maxAgeSecondsStr []byte
var prefix string

const resetScript = `redis.call('DEL', ARGV[1])
redis.call('HSET', ARGV[1], '.created', ARGV[2])
redis.call('EXPIRE', ARGV[1], ARGV[3])`

var resetScriptHash string

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			cli = cfg.GetRedisClient()
			cookieName = cfg.Session.CookieName
			headerName = cfg.Session.HeaderName
			headerExpireName = headerName + "-Expire"
			maxAgeSeconds = int64(cfg.Session.MaxAge)
			maxAge = time.Second * time.Duration(cfg.Session.MaxAge)
			qOfMaxAge = maxAgeSeconds * 3 / 4
			maxAgeSecondsStr = utils.B(strconv.FormatInt(maxAgeSeconds, 10))
			prefix = cfg.Session.StorageKeyPrefix
			var err error
			resetScriptHash, err = cli.ScriptLoad(context.Background(), resetScript).Result()
			if err != nil {
				panic(err)
			}
		},
	)
}

func response(sid []byte, ctx *sha.RequestCtx, byCookie bool) {
	if byCookie {
		ctx.Response.SetCookie(cookieName, utils.S(sid), nil)
	} else {
		ctx.Response.Header.Set(headerName, sid)
		ctx.Response.Header.Set(headerExpireName, maxAgeSecondsStr)
	}
}

type Type string

func New(ctx *sha.RequestCtx) Type {
	v := ctx.Get(internal.UserDataKeys.Session)
	if v != nil {
		return v.(Type)
	}

	var sid []byte
	var key string
	var byCookie bool

	if len(headerName) > 0 {
		sid, _ = ctx.Request.Header.Get(headerName)
	}

	if len(sid) < 1 && len(cookieName) > 0 {
		sid, _ = ctx.Request.Cookie(cookieName)
		byCookie = true
	}

	if len(sid) > 0 {
		key = prefix + utils.S(sid)
		ct, err := cli.HGet(ctx, key, ".created").Int64()
		if err != nil {
			sid = nil
		} else {
			if time.Now().Unix()-ct >= qOfMaxAge {
				cli.HSet(ctx, key, ".created", time.Now().Unix())
				cli.Expire(ctx, key, maxAge)
			}
			return Type(key)
		}
	}

	sid = utils.B(xid.New().String())
	key = prefix + utils.S(sid)
	if err := cli.EvalSha(ctx, resetScriptHash, nil, key, time.Now().Unix(), maxAgeSeconds).Err(); err != nil {
		if err != redis.Nil {
			panic(err)
		}
	}
	ctx.Set(internal.UserDataKeys.Session, Type(key))
	response(sid, ctx, byCookie)
	return Type(key)
}

func (s Type) Get(ctx context.Context, key string, dist interface{}) bool {
	p, e := cli.HGet(ctx, string(s), key).Bytes()
	if e != nil {
		if e == redis.Nil {
			return false
		}
		panic(e)
	}
	return json.Unmarshal(p, dist) == nil
}

func (s Type) Set(ctx context.Context, key string, val interface{}) {
	p, e := json.Marshal(val)
	if e != nil {
		panic(e)
	}
	cli.HSet(ctx, string(s), key, p)
}

func (s Type) SetNX(ctx context.Context, key string, val interface{}) bool {
	p, e := json.Marshal(val)
	if e != nil {
		panic(e)
	}
	ret, err := cli.HSetNX(ctx, string(s), key, p).Result()
	if err != nil {
		panic(err)
	}
	return ret
}

func (s Type) Del(ctx context.Context, keys ...string) { cli.HDel(ctx, string(s), keys...) }

func (s Type) Refresh(ctx context.Context) { cli.Expire(ctx, string(s), maxAge) }

func (s Type) IncrBy(ctx context.Context, key string, increment int64) int64 {
	v, e := cli.HIncrBy(ctx, string(s), key, increment).Result()
	if e != nil {
		panic(e)
	}
	return v
}

func (s Type) Clear(ctx context.Context) {
	cli.EvalSha(ctx, resetScriptHash, nil, string(s), time.Now().Unix(), maxAgeSeconds)
}
