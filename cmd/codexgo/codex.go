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
	"github.com/bastean/codexgo/internal/pkg/service/logger/log"
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

	log.Starting("services")

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started("services")

	log.Starting("server")

	go func() {
		if err := server.Run(port); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started("server")

	log.Info("server listening on :" + port)

	log.Info("press ctrl+c to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping("server")

	errServer := server.Stop(ctx)

	log.Stopped("server")

	log.Stopping("services")

	errService := service.Stop(ctx)

	log.Stopped("services")

	if err := errors.Join(errServer, errService); err != nil {
		log.Error(err.Error())
	}

	<-ctx.Done()

	log.Info("exiting...")
}
