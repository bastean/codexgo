package messages_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/stretchr/testify/suite"
)

type RoutingKeyTestSuite struct {
	suite.Suite
}

func (suite *RoutingKeyTestSuite) SetupTest() {}

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
	components := &messages.RoutingKeyComponents{}
	suite.Panics(func() { messages.NewRoutingKey(components) })
}

func TestUnitRoutingKeySuite(t *testing.T) {
	suite.Run(t, new(RoutingKeyTestSuite))
}
