package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

const (
	App     = "codexgo"
	Version = "4.17.2"
)

var (
	err error
)

var (
	isVersion bool
	isDemo    bool
	envFile   string
)

func usage() {
	log.Logo(Version)

	fmt.Print("Example CRUD project applying Hexagonal Architecture, DDD, EDA, CQRS, BDD, CI, and more... in Go.\n\n")

	fmt.Printf("Usage: %s [flags]\n\n", App)

	flag.PrintDefaults()
}

func Up() error {
	flag.BoolVar(&isVersion, "v", false, "Print version")

	flag.BoolVar(&isDemo, "demo", false, "Use preset ENV values")

	flag.StringVar(&envFile, "env", "", "Path to custom ENV file")

	flag.Usage = usage

	flag.Parse()

	switch {
	case isVersion:
		println(App, Version)
		os.Exit(0)
	case isDemo:
		if err = env.InitDemo(); err != nil {
			return errors.BubbleUp(err)
		}
	case envFile != "":
		if err = godotenv.Load(envFile); err != nil {
			return errors.New[errors.Internal](&errors.Bubble{
				What: "Failure to load ENV file",
				Why: errors.Meta{
					"File": envFile,
				},
				Who: err,
			})
		}
	}

	return nil
}
