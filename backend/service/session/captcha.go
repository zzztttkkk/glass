package session

import (
	"errors"
	"github.com/zzztttkkk/sha"
	"github.com/zzztttkkk/sha/captcha"
	"github.com/zzztttkkk/sha/utils"
	"glass/config"
	"glass/internal"
	"image"
	"math/rand"
	"strings"
	"time"
)

var captchaOptions = &captcha.Options{OffsetX: 12, OffsetY: 12, Points: 240}
var skip bool

func init() {
	internal.DigContainer.Append(
		func(cfg *config.Type) {
			if len(cfg.Session.CaptchaFonts) < 1 {
				panic(errors.New("glass.service.session: empty captcha font list"))
			}

			captcha.Init(cfg.Session.CaptchaFonts...)
			skip = cfg.Session.CaptchaSkip
		},
	)
}

func shuffleOne(v []string) string {
	ind := int(rand.Uint32()) % len(v)
	var buf strings.Builder

	for i, t := range v {
		if i == ind {
			a := []rune(t)
			rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
			for _, r := range a {
				buf.WriteRune(r)
			}
		} else {
			buf.WriteString(t)
		}
	}
	return buf.String()
}

// One picture per minute
func (s Type) CaptchaGenPNG(ctx *sha.RequestCtx) image.Image {
	var lastCTime int64
	var now = time.Now().Unix()
	if s.Get(ctx, internal.SessionKeyCaptchaTime, &lastCTime) && now-lastCTime < 60 {
		panic(sha.StatusError(sha.StatusTooManyRequests))
	}

	poem := RandTangPoem()
	ind := int(rand.Uint32()) % len(poem.Contents)
	verse := poem.Contents[ind]

	s.Set(ctx, internal.SessionKeyCaptchaText, verse)
	s.Set(ctx, internal.SessionKeyCaptchaTime, now)

	return captcha.RenderSomeFonts(-1, shuffleOne(poem.nContents[ind]), captchaOptions)
}

var exr = map[rune]struct{}{}

func init() {
	for _, r := range " ,.!?'\"，。？《》“”~……" {
		exr[r] = struct{}{}
	}
}

func (s Type) CaptchaVerify(ctx *sha.RequestCtx) {
	if skip {
		return
	}

	defer s.Del(ctx, internal.SessionKeyCaptchaText)

	v, _ := ctx.FormValue("captcha")
	if len(v) < 1 {
		panic(sha.StatusError(sha.StatusBadRequest))
	}

	var txt string
	if !s.Get(ctx, internal.SessionKeyCaptchaText, &txt) {
		panic(sha.StatusError(sha.StatusBadRequest))
	}
	var ctime int64
	if !s.Get(ctx, internal.SessionKeyCaptchaTime, &ctime) || time.Now().Unix()-ctime > 300 {
		panic(sha.StatusError(sha.StatusBadRequest))
	}

	var n []rune
	for _, r := range utils.S(v) {
		if _, ok := exr[r]; !ok {
			n = append(n, r)
		}
	}

	i := 0
	l := len(n)
	for _, r := range txt {
		if i >= l {
			panic(sha.StatusError(sha.StatusBadRequest))
		}
		if r == ' ' {
			continue
		}
		//goland:noinspection GoNilness
		if r != n[i] {
			panic(sha.StatusError(sha.StatusBadRequest))
		}
		i++
	}
}
