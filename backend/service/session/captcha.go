package session

import (
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
			captcha.Init(cfg.Session.CaptchaFonts...)
			skip = cfg.Session.CaptchaSkip
		},
	)
}

func shuffleOne(v []string) string {
	var buf strings.Builder
	switch rand.Uint32() & 1 {
	case 0:
		a := []rune(v[0])
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		for _, r := range a {
			buf.WriteRune(r)
		}
		buf.WriteString(v[1])
	case 1:
		buf.WriteString(v[0])
		a := []rune(v[1])
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		for _, r := range a {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// One picture per minute
func (s Type) CaptchaGenPNG(ctx *sha.RequestCtx) image.Image {
	var lastCTime int64
	var now = time.Now().Unix()
	if s.Get(ctx, internal.SessionKeys.CaptchaTime, &lastCTime) && now-lastCTime < 60 {
		panic(sha.StatusError(sha.StatusTooManyRequests))
	}

	poem := RandTangPoem()
	ind := int(rand.Uint32()) % len(poem.Contents)
	verse := poem.Contents[ind]

	s.Set(ctx, internal.SessionKeys.CaptchaVerse, verse)
	s.Set(ctx, internal.SessionKeys.CaptchaTime, now)

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

	defer s.Del(ctx, internal.SessionKeys.CaptchaVerse)

	v, _ := ctx.FormValue("captcha")
	if len(v) < 1 {
		panic(sha.StatusError(sha.StatusBadRequest))
	}

	var txt string
	if !s.Get(ctx, internal.SessionKeys.CaptchaVerse, &txt) {
		panic(sha.StatusError(sha.StatusBadRequest))
	}
	var ctime int64
	if !s.Get(ctx, internal.SessionKeys.CaptchaTime, &ctime) || time.Now().Unix()-ctime > 300 {
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
