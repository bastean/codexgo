package server

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/bastean/codexgo/v4/internal/app/server/router"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
)

var (
	Server = &struct {
		Gin string
	}{
		Gin: log.Server("Gin"),
	}
)

//go:embed static
var Files embed.FS

var App *http.Server

func Up() error {
	log.Starting(Server.Gin)

	App = &http.Server{
		Addr:    ":" + env.ServerGinPort,
		Handler: router.New(&Files),
	}

	log.Started(Server.Gin)

	log.Info(fmt.Sprintf("%s listening on %s", Server.Gin, env.ServerGinURL))

	if proxy, ok := env.HasServerGinProxy(); ok {
		log.Info(fmt.Sprintf("%s proxy listening on %s", Server.Gin, strings.Replace(env.ServerGinURL, env.ServerGinPort, proxy, 1)))
	}

	if err := App.ListenAndServe(); err != nil {
		log.CannotBeStarted(Server.Gin)
		return errors.BubbleUp(err, "Up")
	}

	return nil
}

func Down(ctx context.Context) error {
	log.Stopping(Server.Gin)

	if err := App.Shutdown(ctx); err != nil {
		log.CannotBeStopped(Server.Gin)
		return errors.BubbleUp(err, "Down")
	}

	log.Stopped(Server.Gin)

	return nil
}
