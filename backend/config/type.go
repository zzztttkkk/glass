package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nsqio/go-nsq"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/groupcache"
	"github.com/zzztttkkk/sha/utils"
	"time"
)

const (
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvProduction  = "production"
)

type CacheExpires struct {
	Default   utils.TomlDuration            `json:"default" toml:"default"`
	Missing   utils.TomlDuration            `json:"missing" toml:"missing"`
	Durations map[string]utils.TomlDuration `json:"durations" toml:"durations"`
	Rand      int64                         `json:"rand" toml:"rand"`
}

// just call this function in app startup
func (e *CacheExpires) Get(key string) time.Duration {
	if len(e.Durations) < 1 {
		return e.Default.Duration
	}
	if v, found := e.Durations[key]; found {
		return v.Duration
	}
	return e.Default.Duration
}

type RedisConfig struct {
	Mode  string   `json:"mode" toml:"mode"`
	Nodes []string `json:"nodes" toml:"nodes"`
}

type Type struct {
	Env    string `json:"env" toml:"env"`
	Secret string `json:"secret" toml:"secret"`

	Static struct {
		WebBuildPath string `json:"web_build_path" toml:"web-build-path"`
	} `json:"static" toml:"static"`

	Auth struct {
		CookieName  string             `json:"cookie_name" toml:"cookie-name"`
		HeaderName  string             `json:"header_name" toml:"header-name"`
		TokenMaxAge utils.TomlDuration `json:"token_max_age" toml:"max-age"`
	} `json:"auth" toml:"auth"`

	Cache struct {
		Expires            *groupcache.Expires `json:"expires" toml:"expires"`
		MaxWaitWhenCaching int32               `json:"max_wait_when_caching" toml:"max-wait-when-caching"`
		Redis              *RedisConfig        `json:"redis" toml:"redis"`
	} `json:"cache" toml:"cache"`

	HTTP struct {
		Server            sha.ServerConf            `json:"server" toml:"server"`
		HTTPProtocol      sha.HTTPConf              `json:"http_protocol" toml:"http-protocol"`
		WebSocketProtocol sha.WebSocketProtocolConf `json:"web_socket_protocol" toml:"web-socket-protocol"`
		Router            sha.MuxConf               `json:"router" toml:"router"`

		CORS struct {
			Origins []struct {
				Origin string `json:"origin" toml:"origin"`
				sha.CorsConfig
			} `json:"origins" toml:"origins"`
			Skip bool `json:"skip" toml:"skip"`
		} `json:"cors" toml:"cors"`

		CSRF struct {
			CookieName  string             `json:"cookie_name" toml:"cookie-name"`
			HeaderName  string             `json:"header_name" toml:"header-name"`
			StorageName string             `json:"storage_name" toml:"storage-name"`
			MaxAge      utils.TomlDuration `json:"max_age" toml:"max-age"`
			Skip        bool               `json:"skip" toml:"skip"`
		} `json:"csrf" toml:"csrf"`
	} `json:"http" toml:"http"`

	Session struct {
		CookieName       string             `json:"cookie_name" toml:"cookie-name"`
		HeaderName       string             `json:"header_name" toml:"header-name"`
		MaxAge           utils.TomlDuration `json:"max_age" toml:"max-age"`
		StorageKeyPrefix string             `json:"storage_key_prefix" toml:"storage-key-prefix"`

		CaptchaFonts []string `json:"captcha_fonts" toml:"captcha-fonts"`
		CaptchaSkip  bool     `json:"captcha_skip" toml:"captcha-skip"`
	} `json:"session" toml:"session"`

	Database struct {
		DriverName   string   `json:"driver_name" toml:"driver-name"`
		WriteableURI string   `json:"writeable_uri" toml:"writeable-uri"`
		ReadonlyURIs []string `json:"readonly_uris" toml:"readonly-uris"`
	} `json:"database" toml:"database"`

	Redis RedisConfig `json:"redis" toml:"redis"`

	NSQ struct {
		Addr       string `json:"addr" toml:"addr"`
		AuthSecret string `json:"auth_secret" toml:"auth-secret"`
	} `json:"nsq" toml:"nsq"`
}

var cacheRedisCli redis.Cmdable
var redisCli redis.Cmdable
var nsqProducer *nsq.Producer

func (t *Type) GetRedisClient() redis.Cmdable {
	if redisCli == nil {
		redisCli = initRedisClient(&t.Redis)
	}
	return redisCli
}

func (t *Type) GetCacheStorage() groupcache.Storage {
	if t.Cache.Redis == nil {
		return groupcache.RedisStorage(t.GetRedisClient())
	}
	if cacheRedisCli == nil {
		cacheRedisCli = initRedisClient(t.Cache.Redis)
	}
	return groupcache.RedisStorage(cacheRedisCli)
}

func (t *Type) GetNsqProducer() *nsq.Producer {
	if nsqProducer == nil {
		var err error
		conf := nsq.NewConfig()
		conf.AuthSecret = t.NSQ.AuthSecret
		nsqProducer, err = nsq.NewProducer(t.NSQ.Addr, conf)
		if err != nil {
			panic(err)
		}
	}
	return nsqProducer
}

func (t *Type) NewNsqConsumer(topic string, channel string) *nsq.Consumer {
	conf := nsq.NewConfig()
	conf.AuthSecret = t.NSQ.AuthSecret

	v, e := nsq.NewConsumer(topic, channel, conf)
	if e != nil {
		panic(e)
	}
	return v
}

var allowAnyCORS = sha.NewCorsOptions(
	&sha.CorsConfig{
		AllowMethods: "*", AllowCredentials: true,
		AllowHeaders: "*", ExposeHeaders: "*",
	},
)

func (t *Type) GetCORSChecker() sha.CORSOriginChecker {
	if t.HTTP.CORS.Skip {
		return func(origin []byte) *sha.CorsOptions { return allowAnyCORS }
	}

	var co sha.CORSOriginChecker
	if len(t.HTTP.CORS.Origins) > 0 {
		m := map[string]*sha.CorsOptions{}
		for _, v := range t.HTTP.CORS.Origins {
			m[v.Origin] = sha.NewCorsOptions(&v.CorsConfig)
		}
		co = func(origin []byte) *sha.CorsOptions {
			fmt.Println(string(origin))
			return m[utils.S(origin)]
		}
	}
	return co
}

func Default() Type {
	v := Type{
		Secret: "$ENV{GLASS_SECRET}",
	}

	v.Cache.MaxWaitWhenCaching = 20

	v.Auth.CookieName = "_gac"
	v.Auth.TokenMaxAge = utils.TomlDuration{Duration: time.Hour * 24 * 7}

	v.Session.StorageKeyPrefix = "session:"
	v.Session.CookieName = "_gsc"

	return v
}
