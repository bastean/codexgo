package messages_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type RegExpTestSuite struct {
	suite.Default
}

func (s *RegExpTestSuite) TestSentinel() {
	s.Equal(`([a-z0-9]{1,30})`, messages.RExOrganization)
	s.Equal(`([a-z0-9]{1,30})`, messages.RExService)
	s.Equal(`(\d+)`, messages.RExVersion)
	s.Equal(`(event|command|query|response)`, messages.RExType)
	s.Equal(`([a-z]{1,30})`, messages.RExEntity)
	s.Equal(`([a-z_]{1,30})`, messages.RExTrigger)
	s.Equal(`([a-z]{1,30})`, messages.RExAction)
	s.Equal(`(queued|succeeded|failed|done)`, messages.RExStatus)

	s.Equal("event", messages.Type.Event)
	s.Equal("command", messages.Type.Command)
	s.Equal("query", messages.Type.Query)
	s.Equal("response", messages.Type.Response)

	s.Equal("queued", messages.Status.Queued)
	s.Equal("succeeded", messages.Status.Succeeded)
	s.Equal("failed", messages.Status.Failed)
	s.Equal("done", messages.Status.Done)
}

func TestUnitRegExpSuite(t *testing.T) {
	suite.Run(t, new(RegExpTestSuite))
}
