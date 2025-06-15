package token_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type TokenTestSuite struct {
	suite.Default
}

func (s *TokenTestSuite) TestSentinel() {
	s.Equal(3, token.Limit)
	s.Equal(int(time.Minute*3), token.Every)
	s.Equal(int(time.Hour), token.Next)
}

func (s *TokenTestSuite) TestNew() {
	id := values.Mother().IDValid().Value()

	aggregate, err := token.New(id)

	s.NoError(err)

	s.Equal(aggregate.ID.Value(), id)

	s.Equal(aggregate.Attempt.Limit.Value(), token.Limit)

	s.Equal(aggregate.Attempt.Every.Value(), token.Every)

	s.Equal(aggregate.Attempt.Next.Value(), token.Next)
}

func TestUnitTokenSuite(t *testing.T) {
	suite.Run(t, new(TokenTestSuite))
}
