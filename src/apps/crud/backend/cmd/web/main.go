package main

import (
	"embed"

	"github.com/bastean/codexgo/backend/cmd/web/server"
	"github.com/bastean/codexgo/backend/internal/service"
)

//go:embed static
var Files embed.FS

func main() {
	if err := server.Init(&Files).Run(); err != nil {
		service.Logger.Fatal(err.Error())
	}
}
