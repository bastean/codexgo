package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type RoutingKeyTestSuite struct {
	suite.Suite
}

func (suite *RoutingKeyTestSuite) TestWithValidValue() {
	components := &messages.RoutingKeyComponents{
		Organization: "codexgo",
		Service:      "user",
		Version:      "1",
		Type:         messages.Type.Event,
		Entity:       "user",
		Event:        "created",
		Status:       messages.Status.Succeeded,
	}

	expected := "codexgo.user.1.event.user.created.succeeded"

	actual := messages.NewRoutingKey(components)

	suite.Equal(expected, actual)
}

func (suite *RoutingKeyTestSuite) TestWithInvalidValue() {
	suite.Panics(func() { messages.NewRoutingKey(&messages.RoutingKeyComponents{}) })
}

func TestUnitRoutingKeySuite(t *testing.T) {
	suite.Run(t, new(RoutingKeyTestSuite))
}
