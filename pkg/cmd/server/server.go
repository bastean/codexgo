package server

import (
	"context"
	"embed"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/codexgo/pkg/cmd/server/router"
	"github.com/bastean/codexgo/pkg/cmd/server/service"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
)

//go:embed static
var Files embed.FS

func Run(port string) {
	logger.Info("starting services")

	err := service.Start()

	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("starting server")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router.New(&Files),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	logger.Info("listening and serving HTTP on :" + port)

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	logger.Info("stopping services")

	errService := service.Stop(ctx)

	logger.Info("stopping server")

	errServer := server.Shutdown(ctx)

	err = errors.Join(errService, errServer)

	if err != nil {
		logger.Error(err.Error())
	}

	<-ctx.Done()

	logger.Info("exiting server")
}
