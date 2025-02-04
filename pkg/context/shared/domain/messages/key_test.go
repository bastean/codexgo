package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type KeyTestSuite struct {
	suite.Suite
}

func (s *KeyTestSuite) TestWithValidValue() {
	components := &messages.KeyComponents{
		Organization: "codexgo",
		Service:      "user",
		Version:      "1",
		Type:         messages.Type.Event,
		Entity:       "user",
		Event:        "created",
		Status:       messages.Status.Succeeded,
	}

	expected := messages.Key("codexgo.user.1.event.user.created.succeeded")

	actual := messages.NewKey(components)

	s.Equal(expected, actual)
}

func (s *KeyTestSuite) TestWithInvalidValue() {
	s.Panics(func() { messages.NewKey(new(messages.KeyComponents)) })
}

func TestUnitKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
