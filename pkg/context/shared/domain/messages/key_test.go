package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type KeyTestSuite struct {
	suite.Suite
}

func (s *KeyTestSuite) SetupSuite() {
	s.Equal(messages.RExKeyOrganization, `([a-z0-9]{1,20})`)
	s.Equal(messages.RExKeyService, `([a-z0-9]{1,20})`)
	s.Equal(messages.RExKeyVersion, `(\d+)`)
	s.Equal(messages.RExKeyType, `(event|command|query|response)`)
	s.Equal(messages.RExKeyEntity, `([a-z]{1,20})`)
	s.Equal(messages.RExKeyAction, `([a-z]{1,20})`)
	s.Equal(messages.RExKeyStatus, `(queued|succeeded|failed|done)`)

	s.Equal(messages.RExKeyComponents, `^([a-z0-9]{1,20})\.([a-z0-9]{1,20})\.(\d+)\.(event|command|query|response)\.([a-z]{1,20})\.([a-z]{1,20})\.(queued|succeeded|failed|done)$`)
}

func (s *KeyTestSuite) TestWithValidValue() {
	key, err := values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
		Organization: "codexgo",
		Service:      "user",
		Version:      "1",
		Type:         messages.Type.Event,
		Entity:       "user",
		Action:       "created",
		Status:       messages.Status.Succeeded,
	}))

	s.NoError(err)

	actual := key.Value()

	expected := "codexgo.user.1.event.user.created.succeeded"

	s.Equal(expected, actual)
}

func (s *KeyTestSuite) TestWithInvalidValue() {
	expected := "(Validate): Key has an invalid nomenclature"

	s.PanicsWithValue(expected, func() {
		_, _ = values.New[*messages.Key](messages.ParseKey(new(messages.KeyComponents)))
	})
}

func TestUnitKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
