package cli

import (
	"flag"
	"fmt"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/joho/godotenv"
)

const (
	cli = "codexgo"
)

var (
	env string
)

func usage() {
	log.Logo()

	fmt.Print("Example CRUD project applying Hexagonal Architecture, DDD, EDA, CQRS, BDD, CI, and more... in Go.\n\n")

	fmt.Printf("Usage: %s [flags]\n\n", cli)

	flag.PrintDefaults()
}

func Up() error {
	flag.StringVar(&env, "env", "", "Path to ENV file (required)")

	flag.Usage = usage

	flag.Parse()

	if err := godotenv.Load(env); err != nil && env != "" {
		return errors.NewInternal(&errors.Bubble{
			Where: "Up",
			What:  "Failure to load ENV file",
			Who:   err,
		})
	}

	return nil
}
