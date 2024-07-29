package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/codexgo/v4/internal/app/server"
	"github.com/bastean/codexgo/v4/internal/pkg/service"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
)

const cli = "codexgo"

var (
	err error
)

var (
	Services = "Services"
	Apps     = "Apps"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n\n", cli)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&env.ServerGinPort, "port", env.ServerGinPort, "Gin Server Port (optional)")

	flag.Usage = usage

	flag.Parse()

	log.Logo()

	log.Starting(Services)

	if err = service.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started(Services)

	log.Starting(Apps)

	go func() {
		if err := server.Up(); err != nil {
			log.Fatal(err.Error())
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
