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
	"github.com/bastean/codexgo/pkg/cmd/server/service/broker"
	"github.com/bastean/codexgo/pkg/cmd/server/service/database"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/service/notify"
	"github.com/bastean/codexgo/pkg/cmd/server/service/user"
)

//go:embed static
var Files embed.FS

func Run(port string) {
	errNotify := notify.Init()
	errBroker := broker.Init()
	errDatabase := database.Init()
	errUser := user.Init()

	err := errors.Join(errNotify, errBroker, errDatabase, errUser)

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

	logger.Info("closing server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	errBroker = broker.Close()

	errDatabase = database.Close(ctx)

	errServer := server.Shutdown(ctx)

	err = errors.Join(errBroker, errDatabase, errServer)

	if err != nil {
		logger.Error(err.Error())
	}

	<-ctx.Done()

	logger.Info("exiting server")
}
