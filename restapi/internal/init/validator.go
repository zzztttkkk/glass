package init

import (
	"github.com/zzztttkkk/sha/utils"
	"github.com/zzztttkkk/sha/validator"
	"regexp"
	"strings"
	"unicode"
)

func init() {
	validator.RegisterBytesFilterWithDescription(
		"username",
		func(v []byte) ([]byte, bool) {
			var b strings.Builder
			for _, r := range utils.S(v) {
				if unicode.IsSpace(r) {
					continue
				}
				b.WriteRune(r)
			}
			var ret = b.String()
			if len(ret) < 3 {
				return nil, false
			}
			return utils.B(ret), true
		},
		"remove all space and ensure the remain length > 2",
	)

	validator.RegisterRegexp(
		"password",
		regexp.MustCompile("[a-zA-Z0-9`~!@#$%^&*()-=_+{}\\[\\];':\".<>,/?\\\\|]{6,64}"),
	)
}
