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
	"github.com/bastean/codexgo/internal/pkg/service/env"
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
	flag.StringVar(&port, "p", env.Server.Port, "Port")

	flag.Usage = usage

	flag.Parse()

	log.Starting("services")

	if err := service.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started("services")

	log.Starting("server")

	go func() {
		if err := server.Up(port); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started("server")

	log.Info("server listening on :" + port)

	if proxy, ok := env.Server.HasProxy(); ok {
		log.Info("server proxy listening on :" + proxy)
	}

	log.Info("press ctrl+c to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping("server")

	errServer := server.Down(ctx)

	log.Stopped("server")

	log.Stopping("services")

	errService := service.Down(ctx)

	log.Stopped("services")

	if err := errors.Join(errServer, errService); err != nil {
		log.Error(err.Error())
	}

	<-ctx.Done()

	log.Info("exiting...")
}
