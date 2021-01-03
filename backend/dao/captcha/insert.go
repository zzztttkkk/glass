package captcha

import (
	"bufio"
	"bytes"
	"context"
	"github.com/zzztttkkk/sha/sqlx"
	"os"
	"time"
)

func (Namespace) Create(ctx context.Context, txt string) {
	op.Insert(
		ctx,
		sqlx.Data{
			"created_at": time.Now().UnixNano(),
			"txt":        txt,
		},
	)
}

func (Namespace) Load(ctx context.Context, f *os.File) {
	stmt, err := sqlx.Exe(ctx).Exe.PreparexContext(ctx, insertSql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	now := time.Now().UnixNano()
	var scanner = bufio.NewScanner(f)

	for scanner.Scan() {
		v := scanner.Bytes()
		v = bytes.TrimSpace(v)
		if len(v) < 1 {
			continue
		}
		_, err = stmt.Exec(now, v)
		if err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
