package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/codexgo/internal/app/server"
	"github.com/bastean/codexgo/internal/pkg/service"
	"github.com/bastean/codexgo/internal/pkg/service/logger"
)

const cli = "codexgo"

var port string

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", cli)
	fmt.Printf("\nE.g.: %s -p 8080\n\n", cli)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&port, "p", os.Getenv("CODEXGO_SERVER_GIN_PORT"), "Port")

	flag.Usage = usage

	flag.Parse()

	logger.Starting("services")

	if err := service.Run(); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Started("services")

	logger.Starting("server")

	go func() {
		if err := server.Run(port); err != nil {
			logger.Fatal(err.Error())
		}
	}()

	logger.Started("server")

	logger.Info("server listening on :" + port)

	logger.Info("press ctrl+c to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	logger.Stopping("server")

	errServer := server.Stop(ctx)

	logger.Stopped("server")

	logger.Stopping("services")

	errService := service.Stop(ctx)

	logger.Stopped("services")

	if err := errors.Join(errServer, errService); err != nil {
		logger.Error(err.Error())
	}

	<-ctx.Done()

	logger.Info("exiting...")
}
