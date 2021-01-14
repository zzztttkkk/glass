package user

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/internal"
	"math/big"
	mrandlib "math/rand"
	"time"

	"github.com/zzztttkkk/sha/sqlx"
)

var secret []byte
var source = mrandlib.NewSource(0)
var mrand = mrandlib.New(source)
var pwdHashPool *utils.HashPool

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			secret = []byte(cfg.Secret)
			if len(secret) < 1 {
				panic("empty secret key")
			}
			pwdHashPool = utils.NewHashPoll(sha512.New512_256, secret)
		},
	)

	var max big.Int
	max.SetInt64(time.Now().UnixNano())
	n, _ := rand.Int(rand.Reader, &max)
	mrand.Seed(n.Int64())

}

func genAccountSecret() []byte {
	bytes := make([]byte, 1024)
	mrand.Read(bytes)
	var ret = make([]byte, 2048)
	hex.Encode(ret, bytes)
	return ret
}

func (Namespace) Insert(ctx context.Context, name, pwd []byte) (int64, []byte) {
	accountSecret := genAccountSecret()
	uid := op.Insert(
		ctx,
		sqlx.Data{
			"name":       name,
			"password":   pwdHashPool.Sum(pwd),
			"secret":     accountSecret,
			"created_at": time.Now().UnixNano(),
		},
	)
	return uid, accountSecret
}
