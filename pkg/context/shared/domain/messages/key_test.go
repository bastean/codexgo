package messages_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type KeyTestSuite struct {
	suite.Suite
}

func (s *KeyTestSuite) SetupSuite() {
	s.Equal(messages.RExKey, `^([a-z0-9]{1,30})\.([a-z0-9]{1,30})\.(\d+)\.(event|command|query|response)\.([a-z]{1,30})\.([a-z]{1,30})\.(queued|succeeded|failed|done)$`)
}

func (s *KeyTestSuite) TestWithValidValue() {
	components := messages.Mother.KeyComponentsValid()

	key, err := values.New[*messages.Key](messages.ParseKey(components))

	s.NoError(err)

	actual := key.Value()

	expected := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s",
		components.Organization,
		components.Service,
		components.Version,
		components.Type,
		components.Entity,
		components.Action,
		components.Status,
	)

	s.Equal(expected, actual)
}

func (s *KeyTestSuite) TestWithInvalidValue() {
	expected := "(Validate): Key has an invalid nomenclature"

	s.PanicsWithValue(expected, func() {
		_, _ = values.New[*messages.Key](messages.ParseKey(messages.Mother.KeyComponentsInvalid()))
	})
}

func TestUnitKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
