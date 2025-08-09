package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bastean/codexgo/v4/internal/app/cli"
	"github.com/bastean/codexgo/v4/internal/app/server"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

const (
	Services = "Services"
	Apps     = "Apps"
)

var (
	err error
)

func main() {
	if err = cli.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Logo()

	log.Starting(Services)

	if err = service.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started(Services)

	log.Starting(Apps)

	go func() {
		if errServer := server.Up(); errServer != nil {
			log.Fatal(errServer.Error())
		}
	}()

	log.Started(Apps)

	log.Info("Press Ctrl+C to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping(Apps)

	if err = server.Down(ctx); err != nil {
		log.Error(err.Error())
	}

	log.Stopped(Apps)

	log.Stopping(Services)

	if err = service.Down(ctx); err != nil {
		log.Error(err.Error())
	}

	log.Stopped(Services)

	<-ctx.Done()

	log.Info("Exiting...")
}
