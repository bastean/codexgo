package embed_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/embed"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type RegExpTestSuite struct {
	suite.Default
}

func (s *RegExpTestSuite) TestSentinel() {
	s.Equal(`\[.*]`, embed.RExEmbed)
}

func TestUnitRegExpSuite(t *testing.T) {
	suite.Run(t, new(RegExpTestSuite))
}
