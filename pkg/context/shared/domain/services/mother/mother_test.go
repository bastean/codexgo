package mother_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
)

type MotherTestSuite struct {
	suite.Suite
}

func (s *MotherTestSuite) SetupSuite() {
	if _, ok := os.LookupEnv("GOTEST"); !ok {
		panic("\"Test Environment (GOTEST)\" not defined")
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
