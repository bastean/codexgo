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
	s.Equal(messages.RExOrganization, `([a-z0-9]{1,30})`)
	s.Equal(messages.RExService, `([a-z0-9]{1,30})`)
	s.Equal(messages.RExVersion, `(\d+)`)
	s.Equal(messages.RExType, `(event|command|query|response)`)
	s.Equal(messages.RExEntity, `([a-z]{1,30})`)
	s.Equal(messages.RExTrigger, `([a-z_]{1,30})`)
	s.Equal(messages.RExAction, `([a-z]{1,30})`)
	s.Equal(messages.RExStatus, `(queued|succeeded|failed|done)`)

	s.Equal(messages.Type.Event, "event")
	s.Equal(messages.Type.Command, "command")
	s.Equal(messages.Type.Query, "query")
	s.Equal(messages.Type.Response, "response")

	s.Equal(messages.Status.Queued, "queued")
	s.Equal(messages.Status.Succeeded, "succeeded")
	s.Equal(messages.Status.Failed, "failed")
	s.Equal(messages.Status.Done, "done")
}

func TestUnitRegExpSuite(t *testing.T) {
	suite.Run(t, new(RegExpTestSuite))
}
