package cli

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

const (
	cli = "codexgo"
)

var (
	err     error
	isDemo  bool
	envFile string
)

func usage() {
	log.Logo()

	fmt.Print("Example CRUD project applying Hexagonal Architecture, DDD, EDA, CQRS, BDD, CI, and more... in Go.\n\n")

	fmt.Printf("Usage: %s [flags]\n\n", cli)

	flag.PrintDefaults()
}

func Up() error {
	flag.BoolVar(&isDemo, "demo", false, "Use preset ENV values")

	flag.StringVar(&envFile, "env", "", "Path to custom ENV file")

	flag.Usage = usage

	flag.Parse()

	switch {
	case isDemo:
		if err = env.InitDemo(); err != nil {
			return errors.BubbleUp(err, "Up")
		}
	case envFile != "":
		if err = godotenv.Load(envFile); err != nil {
			return errors.New[errors.Internal](&errors.Bubble{
				Where: "Up",
				What:  "Failure to load ENV file",
				Why: errors.Meta{
					"File": envFile,
				},
				Who: err,
			})
		}
	}

	return nil
}
