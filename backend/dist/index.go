package dist

import "embed"

//go:embed 300_tang_poems.json
var TangPoems []byte

//go:embed favicon.ico
var Favicon embed.FS
