package choose_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/choose"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type ChooseTestSuite struct {
	suite.Default
}

func (s *ChooseTestSuite) TestOneWithTrue() {
	expected := choose.Mother().LoremIpsumWord()

	actual := choose.One(true, expected, choose.Mother().LoremIpsumWord())

	s.Equal(expected, actual)
}

func (s *ChooseTestSuite) TestOneWithFalse() {
	expected := choose.Mother().LoremIpsumWord()

	actual := choose.One(false, choose.Mother().LoremIpsumWord(), expected)

	s.Equal(expected, actual)
}

func TestUnitChooseSuite(t *testing.T) {
	suite.Run(t, new(ChooseTestSuite))
}
