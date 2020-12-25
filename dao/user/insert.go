package user

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"github.com/zzztttkkk/sha/sqlx"
	"glass/config"
	"glass/internal"
	"hash"
	"math/big"
	mrandlib "math/rand"
	"sync"
	"time"
)

var secret []byte
var source = mrandlib.NewSource(0)
var mrand = mrandlib.New(source)

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			secret = []byte(cfg.Secret)
			if len(secret) < 1 {
				panic("empty secret key")
			}
		},
	)

	var max big.Int
	max.SetInt64(time.Now().UnixNano())
	n, _ := rand.Int(rand.Reader, &max)
	mrand.Seed(n.Int64())
}

var hashPoll = sync.Pool{New: func() interface{} { return nil }}

type _Hash struct {
	hash.Hash
	buf []byte
}

func (i *_Hash) Sum() []byte {
	return i.Hash.Sum(i.buf)
}

func (i *_Hash) reset() *_Hash {
	i.Hash.Reset()
	i.buf = i.buf[:0]
	return i
}

func genPwdHash(pwd []byte) []byte {
	var hw *_Hash
	hi := hashPoll.Get()
	if hi == nil {
		hw = &_Hash{Hash: sha512.New512_256(), buf: make([]byte, 0, 128)}
	} else {
		hw = hi.(*_Hash)
	}
	defer func() { hashPoll.Put(hw.reset()) }()

	_, _ = hw.Write(pwd)
	_, _ = hw.Write(secret)

	ret := make([]byte, 64)
	hex.Encode(ret, hw.Sum())
	return ret
}

func genSecret() []byte {
	bytes := make([]byte, 128)
	mrand.Read(bytes)
	var ret = make([]byte, 256)
	hex.Encode(ret, bytes)
	return ret
}

func (Namespace) Insert(ctx context.Context, name, pwd []byte) (int64, []byte) {
	secret := genSecret()
	uid := op.Insert(
		ctx,
		sqlx.Data{
			"name":       name,
			"password":   genPwdHash(pwd),
			"secret":     secret,
			"created_at": time.Now().UnixNano(),
		},
	)
	return uid, secret
}
