package messages_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type KeyTestSuite struct {
	suite.Default
}

func (s *KeyTestSuite) SetupSuite() {
	s.Equal(messages.RExKey, `^([a-z0-9]{1,30})\.([a-z0-9]{1,30})\.(\d+)\.(event|command|query|response)\.([a-z]{1,30})\.([a-z]{1,30})\.(queued|succeeded|failed|done)$`)
}

func (s *KeyTestSuite) TestWithValidValue() {
	components := messages.Mother().KeyComponentsValid()

	actual := messages.Mother().KeyValidWithComponents(components).Value()

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
	components := messages.Mother().KeyComponentsInvalid()

	expected := fmt.Sprintf("(Validate): Key has an invalid nomenclature %q", messages.ParseKey(components))

	s.PanicsWithValue(expected, func() { messages.Mother().KeyInvalidWithComponents(components) })
}

func TestUnitKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
