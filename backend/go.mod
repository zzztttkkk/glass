module glass

go 1.16

require (
	github.com/go-redis/redis/v8 v8.4.11
	github.com/go-sql-driver/mysql v1.4.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/nsqio/go-nsq v1.0.8 // indirect
	github.com/rs/xid v1.2.1
	github.com/zzztttkkk/sha v0.0.2
)

replace github.com/zzztttkkk/sha v0.0.2 => ../../sha
