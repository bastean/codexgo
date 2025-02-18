package server

import (
	"context"
	"embed"
	"fmt"
	"net/http"

	"github.com/bastean/codexgo/v4/internal/app/server/router"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

var Server = &struct {
	Gin string
}{
	Gin: log.Server("Gin"),
}

var (
	//go:embed static
	Files embed.FS
	App   *http.Server
)

func Up() error {
	log.Starting(Server.Gin)

	App = &http.Server{
		Addr:    ":" + env.ServerGinPort,
		Handler: router.New(&Files),
	}

	log.Started(Server.Gin)

	log.Info(fmt.Sprintf("%s listening on %s", Server.Gin, env.ServerGinURL))

	if err := App.ListenAndServe(); errors.IsNot(err, http.ErrServerClosed) {
		log.CannotBeStarted(Server.Gin)

		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Up",
			What:  "Failure to start Server",
			Who:   err,
		})
	}

	return nil
}

func Down(ctx context.Context) error {
	log.Stopping(Server.Gin)

	if err := App.Shutdown(ctx); err != nil {
		log.CannotBeStopped(Server.Gin)

		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Down",
			What:  "Failure to shutdown Server",
			Who:   err,
		})
	}

	log.Stopped(Server.Gin)

	return nil
}
