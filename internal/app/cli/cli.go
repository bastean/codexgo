package cli

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
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
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Up",
			What:  "Failure to load ENV file",
			Who:   err,
		})
	}

	return nil
}
