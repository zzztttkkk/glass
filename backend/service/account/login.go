package account

import (
	"context"
	"crypto/sha256"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/auth"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/dao"
	"glass/internal"
)

var secret []byte
var authMaxAge int64
var authCookieName string
var authHeaderName string
var authTokenGenerator utils.IDTokenGenerator

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			secret = []byte(cfg.Secret)
			authMaxAge = int64(cfg.Auth.TokenMaxAge)
			authCookieName = cfg.Auth.CookieName
			authHeaderName = cfg.Auth.HeaderName
			authTokenGenerator = utils.NewIDTokenGenerator(utils.NewHashPoll(sha256.New, secret))
		},
	)
}

func (Namespace) DoLogin(ctx context.Context, name, password []byte) (ret string) {
	uid := dao.User.Auth(ctx, name, password)
	if uid < 1 {
		return ""
	}
	return authTokenGenerator.EncodeID(uid, authMaxAge)
}

func (Namespace) Auth(ctx context.Context) (auth.Subject, error) {
	rctx := sha.MustToRCtx(ctx)
	subject := rctx.Get(internal.UserDataKeys.Subject)
	if subject != nil {
		return subject.(auth.Subject), nil
	}

	v, ok := rctx.Request.Cookie(authCookieName)
	if !ok {
		v, ok = rctx.Request.Header.Get(authHeaderName)
	}
	if len(v) < 1 {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	uid, err := authTokenGenerator.DecodeID(utils.S(v))
	if err != nil {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}
	user := dao.User.InfoByID(ctx, uid)
	if user == nil {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	rctx.Set(internal.UserDataKeys.Subject, user)
	return user, nil
}
