package suite

import (
	"os"
	"strconv"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

var (
	Run = suite.Run
)

type Default struct {
	suite.Suite
}

type Frozen struct {
	suite.Suite
}

func (s *Frozen) SetupTest() {
	s.NoError(os.Setenv("GOTEST_FROZEN", "1"))
}

func (s *Frozen) SetTimeBefore(d time.Duration) {
	s.NoError(os.Setenv("GOTEST_FROZEN_BEFORE", strconv.Itoa(int(d))))
}

func (s *Frozen) UnsetTimeBefore() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN_BEFORE"))
}

func (s *Frozen) SetTimeAfter(d time.Duration) {
	s.NoError(os.Setenv("GOTEST_FROZEN_AFTER", strconv.Itoa(int(d))))
}

func (s *Frozen) UnsetTimeAfter() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN_AFTER"))
}

func (s *Frozen) TearDownTest() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN"))
	s.UnsetTimeBefore()
	s.UnsetTimeAfter()
}
