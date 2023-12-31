package main

import (
	"embed"

	"github.com/bastean/codexgo/backend/cmd/web/server"
)

//go:embed templates
var files embed.FS

func main() {
	server.Init(&files).Run()
}
