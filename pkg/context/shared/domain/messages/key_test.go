package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type KeyTestSuite struct {
	suite.Suite
}

func (suite *KeyTestSuite) TestWithValidValue() {
	components := &messages.KeyComponents{
		Organization: "codexgo",
		Service:      "user",
		Version:      "1",
		Type:         messages.Type.Event,
		Entity:       "user",
		Event:        "created",
		Status:       messages.Status.Succeeded,
	}

	expected := events.Key("codexgo.user.1.event.user.created.succeeded")

	actual := messages.NewKey(components)

	suite.Equal(expected, actual)
}

func (suite *KeyTestSuite) TestWithInvalidValue() {
	suite.Panics(func() { messages.NewKey(&messages.KeyComponents{}) })
}

func TestUnitKeySuite(t *testing.T) {
	suite.Run(t, new(KeyTestSuite))
}
