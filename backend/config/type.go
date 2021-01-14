package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/zzztttkkk/sha"
)

type Type struct {
	Secret string `json:"secret" toml:"secret"`

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

		CorsOptions sha.CorsOptions `json:"cors" toml:"cors"`

		Session struct {
			CookieName       string `json:"cookie_name" toml:"cookie-name"`
			HeaderName       string `json:"header_name" toml:"header-name"`
			MaxAge           int    `json:"max_age" toml:"max-age"`
			StorageKeyPrefix string `json:"storage_key_prefix" toml:"storage-key-prefix"`

			CRSF struct {
				CookieName  string `json:"cookie_name" toml:"cookie-name"`
				HeaderName  string `json:"header_name" toml:"header-name"`
				StorageName string `json:"storage_name" toml:"storage-name"`
				MaxAge      int    `json:"max_age" toml:"max-age"`
			} `json:"crsf" toml:"crsf"`

			CaptchaFonts []string `json:"captcha_fonts" toml:"captcha-fonts"`
		} `json:"session" toml:"session"`
	} `json:"http" toml:"http"`

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

func _Default() Type {
	v := Type{
		Secret: "$ENV{GLASS_SECRET}",
	}
	v.HTTP.Session.StorageKeyPrefix = "session:"
	v.HTTP.Session.MaxAge = 1800
	v.HTTP.Session.HeaderName = "GlassSession"
	v.HTTP.Session.CookieName = "_gsi"
	v.HTTP.Session.CRSF.CookieName = "_gcrsfp"
	v.HTTP.Session.CRSF.HeaderName = "GlassCRSF"
	v.HTTP.Session.CRSF.MaxAge = 900
	return v
}
