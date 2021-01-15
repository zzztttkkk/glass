package internal

var SessionKeys = struct {
	CaptchaVerse string
	CaptchaTime  string
	CaptchaRetry string
}{
	CaptchaTime:  ".c.t",
	CaptchaVerse: ".c.v",
	CaptchaRetry: ".c.r",
}
