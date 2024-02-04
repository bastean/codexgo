package main

import (
	"embed"

	"github.com/bastean/codexgo/backend/cmd/web/server"
	"github.com/bastean/codexgo/backend/internal/container"
)

//go:embed static templates
var Files embed.FS

func main() {
	if err := server.Init(&Files).Run(); err != nil {
		container.Logger.Fatal(err.Error())
	}
}
