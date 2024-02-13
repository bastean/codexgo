package main

import (
	"context"
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/codexgo/backend/cmd/web/router"
	"github.com/bastean/codexgo/backend/internal/service"
)

//go:embed static
var Files embed.FS

var Port = os.Getenv("PORT")
var Server = &http.Server{Addr: ":" + Port, Handler: router.New(&Files)}

func main() {
	service.Logger.Info("starting server")

	go func() {
		if err := Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			service.Logger.Fatal(err.Error())
		}
	}()

	service.Logger.Info("listening and serving HTTP on :" + Port)

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	service.Logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := Server.Shutdown(ctx); err != nil {
		service.Logger.Fatal(err.Error())
	}

	<-ctx.Done()

	service.Logger.Info("server exiting")
}
