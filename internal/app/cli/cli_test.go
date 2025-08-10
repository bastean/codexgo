package cli_test

import (
	"os/exec"
	"testing"

	"github.com/bastean/codexgo/v4/internal/app/cli"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type CLITestSuite struct {
	suite.Default
	SUT string
}

func (s *CLITestSuite) SetupSuite() {
	s.SUT = "../../../cmd/codexgo"
}

func (s *CLITestSuite) TestSentinel() {
	s.Equal("codexgo", cli.App)
	s.Equal("4.16.2", cli.Version)
}

func (s *CLITestSuite) TestHelp() {
	expected := `              _________               ________________ 
_____________ ______  /_____ ____  __ __  ____/__  __ \
_  ___/_  __ \_  __  / _  _ \__  |/_/ _  / __  _  / / /
/ /__  / /_/ // /_/ /  /  __/__>  <   / /_/ /  / /_/ / 
\___/  \____/ \__,_/   \___/ /_/|_|   \____/   \____/  v4.16.2

Example CRUD project applying Hexagonal Architecture, DDD, EDA, CQRS, BDD, CI, and more... in Go.

Usage: codexgo [flags]

  -demo
    	Use preset ENV values
  -env string
    	Path to custom ENV file
  -v	Print version
`

	actual, err := exec.Command("go", "run", s.SUT, "-h").CombinedOutput() //nolint:gosec

	s.NoError(err)

	s.Equal(expected, string(actual))
}

func TestAcceptanceCLISuite(t *testing.T) {
	suite.Run(t, new(CLITestSuite))
}
