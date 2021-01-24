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
	"time"
)

var secret []byte
var authMaxAge int64
var authCookieName string
var authHeaderName string
var AuthTokenGenerator utils.IDTokenGenerator

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			secret = []byte(cfg.Secret)
			authMaxAge = int64(cfg.Auth.TokenMaxAge)
			authCookieName = cfg.Auth.CookieName
			authHeaderName = cfg.Auth.HeaderName
			AuthTokenGenerator = utils.NewIDTokenGenerator(utils.NewHashPoll(sha256.New, secret))
		},
	)
}

func (Namespace) DoLogin(ctx context.Context, name, password []byte) (ret string) {
	uid := dao.User.Auth(ctx, name, password)
	if uid < 1 {
		return ""
	}
	return AuthTokenGenerator.EncodeID(uid)
}

func (Namespace) Auth(ctx context.Context) (auth.Subject, error) {
	rctx := sha.MustToRCtx(ctx)
	si := rctx.Get(internal.UserDataKeySubject)
	if si != nil {
		return si.(auth.Subject), nil
	}

	authByCookie := true
	v, ok := rctx.Request.Cookie(authCookieName)
	if !ok {
		authByCookie = false
		v, ok = rctx.Request.Header.Get(authHeaderName)
	}
	if len(v) < 1 {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	uid, ll, err := AuthTokenGenerator.DecodeID(utils.S(v))
	if err != nil {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	du := time.Now().Unix() - ll
	if du > authMaxAge {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	if du > authMaxAge/2 {
		token := AuthTokenGenerator.EncodeID(si.(auth.Subject).GetID())
		if authByCookie {
			rctx.Response.SetCookie(authCookieName, token, &sha.CookieOptions{MaxAge: authMaxAge})
		} else {
			rctx.Response.Header.Set(authHeaderName, utils.B(token))
		}
	}
	user := dao.User.GetByID(ctx, uid)
	if user == nil {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	rctx.Set(internal.UserDataLastLogin, ll)
	rctx.Set(internal.UserDataKeySubject, user)
	return user, nil
}
