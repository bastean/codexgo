package server

import (
	"context"
	"embed"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/codexgo/pkg/cmd/server/router"
	_ "github.com/bastean/codexgo/pkg/cmd/server/service/broker"
	_ "github.com/bastean/codexgo/pkg/cmd/server/service/database"
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
)

//go:embed static
var Files embed.FS

func Run(port string) {
	logger.Logger.Info("starting server")

	server := &http.Server{Addr: ":" + port, Handler: router.New(&Files)}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal(err.Error())
		}
	}()

	logger.Logger.Info("listening and serving HTTP on :" + port)

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	logger.Logger.Info("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Logger.Fatal(err.Error())
	}

	<-ctx.Done()

	logger.Logger.Info("server exiting")
}
