module glass

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.4.4
	github.com/go-sql-driver/mysql v1.4.0
	github.com/imdario/mergo v0.3.11
	github.com/jmoiron/sqlx v1.2.0
	github.com/rs/xid v1.2.1
	github.com/zzztttkkk/sha v0.0.2
)

replace github.com/zzztttkkk/sha v0.0.2 => ../../sha
