package account

import (
	"context"
	"github.com/dgrijalva/jwt-go"
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

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			secret = []byte(cfg.Secret)
			authMaxAge = int64(cfg.AuthMaxAge)
			authCookieName = cfg.HTTP.AuthCookieName
			authHeaderName = cfg.HTTP.AuthHeaderName
		},
	)
}

type AuthTokenClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"id"`
}

func (Namespace) DoLogin(ctx context.Context, name, password []byte) (ret string) {
	uid := dao.User.Auth(ctx, name, password)
	if uid < 1 {
		return ""
	}

	var claims AuthTokenClaims
	claims.UserID = uid
	claims.ExpiresAt = time.Now().Unix() + authMaxAge

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &claims)
	ret, _ = token.SignedString(secret)
	return ret
}

func parseToken(ctx context.Context, v []byte) (auth.Subject, error) {
	if len(v) < 1 {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	var claims AuthTokenClaims
	token, err := jwt.ParseWithClaims(utils.S(v), &claims, func(token *jwt.Token) (interface{}, error) { return secret, nil })
	if err != nil || !token.Valid {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}

	user := dao.User.InfoByID(ctx, claims.UserID)
	if user == nil {
		return nil, sha.StatusError(sha.StatusUnauthorized)
	}
	return user, nil
}

func (Namespace) Auth(ctx context.Context) (auth.Subject, error) {
	rctx := sha.RCtx(ctx)
	req := &rctx.Request

	var token []byte

	// http basic authorization
	authorization, _ := req.Header.Get(sha.HeaderAuthorization)
	if len(authorization) > 7 {
		return parseToken(ctx, authorization[7:])
	}

	// header
	if len(token) < 1 && len(authHeaderName) > 0 {
		token, _ = req.Header.Get(authHeaderName)
		return parseToken(ctx, token)
	}

	// cookie
	if len(token) < 1 && len(authHeaderName) > 0 {
		token, _ = req.Cookie(authCookieName)
		return parseToken(ctx, token)
	}
	return nil, sha.StatusError(sha.StatusUnauthorized)
}
