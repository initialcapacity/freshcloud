package freshctl

import (
	"embed"
	_ "embed"
)

var (
	//go:embed resources
	Resources embed.FS
)
