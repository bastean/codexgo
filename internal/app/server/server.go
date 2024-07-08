package server

import (
	"context"
	"embed"
	"net/http"

	"github.com/bastean/codexgo/internal/app/server/router"
)

//go:embed static
var Files embed.FS

var Server *http.Server

func Run(port string) error {
	Server := &http.Server{
		Addr:    ":" + port,
		Handler: router.New(&Files),
	}

	return Server.ListenAndServe()
}

func Stop(ctx context.Context) error {
	return Server.Shutdown(ctx)
}
