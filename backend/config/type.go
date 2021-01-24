package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/utils"
)

const (
	EnvDevelopment = "development"
	EnvTesting     = "testing"
	EnvProduction  = "production"
)

type Type struct {
	Env    string `json:"env" toml:"env"`
	Secret string `json:"secret" toml:"secret"`

	Static struct {
		WebBuildPath string `json:"web_build_path" toml:"web-build-path"`
	} `json:"static" toml:"static"`

	Auth struct {
		CookieName  string `json:"cookie_name" toml:"cookie-name"`
		HeaderName  string `json:"header_name" toml:"header-name"`
		TokenMaxAge int    `json:"token_max_age" toml:"max-age"`
	} `json:"auth" toml:"auth"`

	HTTP struct {
		// server
		Host                          string `json:"host" toml:"host"`
		Port                          int    `json:"port" toml:"port"`
		TLSCert                       string `json:"tls_cert" toml:"tls-cert"`
		TLSKey                        string `json:"tls_key" toml:"tls-key"`
		MaxConnectionKeepAliveSeconds int    `json:"max_connection_keep_alive_seconds" toml:"max-connection-keep-alive-seconds"`
		ReadTimeoutSeconds            int    `json:"read_timeout_seconds" toml:"read-timeout-seconds"`
		IdleTimeoutSeconds            int    `json:"idle_timeout_seconds" toml:"idle-timeout-seconds"`
		WriteTimeoutSeconds           int    `json:"write_timeout_seconds" toml:"write-timeout-seconds"`
		AutoCompression               bool   `json:"auto_compression" toml:"auto-compression"`

		// http protocol
		MaxRequestFirstLineSize       int `json:"max_request_first_line_size" toml:"max-request-first-line-size"`
		MaxRequestHeaderPartSize      int `json:"max_request_header_part_size" toml:"max-request-header-part-size"`
		MaxRequestBodySize            int `json:"max_request_body_size" toml:"max-request-body-size"`
		ReadBufferSize                int `json:"read_buffer_size" toml:"read-buffer-size"`
		MaxReadBufferSize             int `json:"max_read_buffer_size" toml:"max-read-buffer-size"`
		MaxResponseBodyBufferSize     int `json:"max_response_body_buffer_size" toml:"max-response-body-buffer-size"`
		DefaultResponseSendBufferSize int `json:"default_response_send_buffer_size" toml:"default-response-send-buffer-size"`

		// router
		PathPrefix string `json:"path_prefix" toml:"path-prefix"`

		CORS []struct {
			Origin string `json:"origin" toml:"origin"`
			sha.CorsConfig
		} `json:"cors" toml:"cors"`

		CSRF struct {
			CookieName  string `json:"cookie_name" toml:"cookie-name"`
			HeaderName  string `json:"header_name" toml:"header-name"`
			StorageName string `json:"storage_name" toml:"storage-name"`
			MaxAge      int    `json:"max_age" toml:"max-age"`
		} `json:"csrf" toml:"csrf"`
	} `json:"http" toml:"http"`

	Session struct {
		CookieName       string `json:"cookie_name" toml:"cookie-name"`
		HeaderName       string `json:"header_name" toml:"header-name"`
		MaxAge           int    `json:"max_age" toml:"max-age"`
		StorageKeyPrefix string `json:"storage_key_prefix" toml:"storage-key-prefix"`

		CaptchaFonts []string `json:"captcha_fonts" toml:"captcha-fonts"`
		CaptchaSkip  bool     `json:"captcha_skip" toml:"captcha-skip"`
	} `json:"session" toml:"session"`

	Database struct {
		DriverName   string   `json:"driver_name" toml:"driver-name"`
		WriteableURI string   `json:"writeable_uri" toml:"writeable-uri"`
		ReadonlyURIs []string `json:"readonly_uris" toml:"readonly-uris"`
	} `json:"database" toml:"database"`

	Redis struct {
		Mode  string
		Nodes []string
	}

	internal struct {
		redisCli redis.Cmdable
	}
}

func (t *Type) GetRedisClient() redis.Cmdable {
	if t.internal.redisCli == nil {
		t.initRedisClient()
	}
	return t.internal.redisCli
}

func (t *Type) GetCORSChecker() sha.CORSOriginChecker {
	var co sha.CORSOriginChecker
	if len(t.HTTP.CORS) > 0 {
		m := map[string]*sha.CorsOptions{}
		for _, v := range t.HTTP.CORS {
			m[v.Origin] = sha.NewCorsOptions(&v.CorsConfig)
		}
		co = func(origin []byte) *sha.CorsOptions {
			fmt.Println(string(origin))
			return m[utils.S(origin)]
		}
	}
	return co
}

func _Default() Type {
	v := Type{
		Secret: "$ENV{GLASS_SECRET}",
	}
	v.Auth.CookieName = "_gak"
	v.Auth.HeaderName = "x-Glass-Auth"
	v.Session.StorageKeyPrefix = "session:"
	v.Session.MaxAge = 1800
	v.Session.HeaderName = "X-Glass-Session"
	v.Session.CookieName = "_gsi"
	v.HTTP.CSRF.CookieName = "_gcsrf"
	v.HTTP.CSRF.HeaderName = "X-Glass-CSRF"
	v.HTTP.CSRF.MaxAge = 900
	return v
}
