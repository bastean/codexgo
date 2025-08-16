package embed_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/embed"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type ExtractTestSuite struct {
	suite.Default
}

func (s *ExtractTestSuite) TestExtract() {
	message, expected := embed.Mother().EmbedValid()

	actual := embed.Extract(message)

	s.Equal(expected, actual)
}

func (s *ExtractTestSuite) TestExtractErrMissingEmbed() {
	message := embed.Mother().EmbedInvalid()

	actual := embed.Extract(message)

	expected := ""

	s.Equal(expected, actual)
}

func TestUnitExtractSuite(t *testing.T) {
	suite.Run(t, new(ExtractTestSuite))
}
