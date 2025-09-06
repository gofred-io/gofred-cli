package embed

import (
	"embed"
)

//go:embed web
var web embed.FS

func Web() embed.FS {
	return web
}
