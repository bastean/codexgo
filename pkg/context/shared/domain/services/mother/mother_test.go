package mother_test

import (
	"log"
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type MotherTestSuite struct {
	suite.Default
}

func (s *MotherTestSuite) SetupSuite() {
	if _, ok := os.LookupEnv("GOTEST"); !ok {
		log.Panic("\"Test Environment (GOTEST)\" not defined")
	}
}

func (s *MotherTestSuite) TestUnsetENV() {
	s.NoError(os.Unsetenv("GOTEST"))

	expected := "Use \"Mother\" only in a \"Test Environment\""

	s.PanicsWithValue(expected, func() { _ = mother.New[struct{ *mother.Mother }]() })
}

func TestUnitMotherSuite(t *testing.T) {
	suite.Run(t, new(MotherTestSuite))
}
