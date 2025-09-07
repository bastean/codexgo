package ascii_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/ascii"
)

type ASCIITestSuite struct {
	suite.Default
}

func (s *ASCIITestSuite) TestFixWidth() {
	drawing, expected := ascii.Mother().DrawingValid()

	ascii.FixWidth(drawing)

	for _, actual := range drawing {
		s.Equal(expected, len(actual))
	}
}

func TestUnitASCIISuite(t *testing.T) {
	suite.Run(t, new(ASCIITestSuite))
}
