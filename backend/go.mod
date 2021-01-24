module glass

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/go-redis/redis/v8 v8.4.4
	github.com/go-sql-driver/mysql v1.4.0
	github.com/imdario/mergo v0.3.11
	github.com/jmoiron/sqlx v1.2.0
	github.com/rs/xid v1.2.1
	github.com/zzztttkkk/sha v0.0.2
	golang.org/x/net v0.0.0-20201202161906-c7110b5ffcbb
	golang.org/x/text v0.3.3
)

replace github.com/zzztttkkk/sha v0.0.2 => ../../sha
