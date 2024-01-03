package main

import (
	"embed"

	"github.com/bastean/codexgo/backend/cmd/web/server"
)

//go:embed static templates
var Files embed.FS

func main() {
	server.Init(&Files).Run()
}
