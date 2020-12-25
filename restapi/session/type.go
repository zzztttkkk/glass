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
var maxAgeSecondsStr []byte
var prefix string

const clearScript = `redis.call('DEL', KEYS[1])
redis.call('HSET', KEYS[1], '.created', ARGV[1])
redis.call('EXPIRE', KEYS[1], ARGV[2])`

var clearScriptSha1 string

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			cli = cfg.GetRedisClient()
			cookieName = cfg.HTTP.Session.CookieName
			headerName = cfg.HTTP.Session.HeaderName
			headerExpireName = headerName + "-Expire"
			maxAgeSeconds = int64(cfg.HTTP.Session.MaxAge)
			maxAge = time.Second * time.Duration(cfg.HTTP.Session.MaxAge)
			maxAgeSecondsStr = utils.B(strconv.FormatInt(maxAgeSeconds, 10))
			prefix = cfg.HTTP.Session.StorageKeyPrefix
			var err error
			clearScriptSha1, err = cli.ScriptLoad(context.Background(), clearScript).Result()
			if err != nil {
				panic(err)
			}
		},
	)
}

type Type string

const sessionKey = ".session"

func New(ctx *sha.RequestCtx) Type {
	v := ctx.Get(sessionKey)
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

	c := ctx.Context()

	if len(sid) > 0 {
		key = prefix + utils.S(sid)
		if cli.Exists(c, key).Val() < 1 {
			sid = nil
		} else {
			cli.Expire(c, key, maxAge)
			ctx.Set(sessionKey, Type(key))
			return Type(key)
		}
	}

	key = prefix + xid.New().String()
	cli.EvalSha(c, clearScriptSha1, []string{key}, time.Now().Unix(), maxAge)

	if byCookie {
		ctx.Response.SetCookie(cookieName, key, sha.CookieOptions{MaxAge: maxAgeSeconds})
	} else {
		ctx.Response.Header.Set(headerName, utils.B(key))
		ctx.Response.Header.Set(headerExpireName, maxAgeSecondsStr)
	}
	ctx.Set(sessionKey, Type(key))
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

func (s Type) Del(ctx context.Context, keys ...string) { cli.HDel(ctx, string(s), keys...) }

func (s Type) Refresh(ctx context.Context) { cli.Expire(ctx, string(s), maxAge) }

func (s Type) Clear(ctx context.Context) {
	cli.EvalSha(ctx, clearScriptSha1, []string{string(s)}, time.Now().Unix(), maxAge)
}