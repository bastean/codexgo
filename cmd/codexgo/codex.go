package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server"
)

const cli = "codexgo"

var port string

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", cli)
	fmt.Printf("\nE.g.: %s -p 8080\n\n", cli)
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&port, "p", os.Getenv("PORT"), "Port")

	flag.Usage = usage

	flag.Parse()

	server.Run(port)
}
