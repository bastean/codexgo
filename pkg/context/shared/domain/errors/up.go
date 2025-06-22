package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func BubbleUp(who error) error {
	where, _, _, _ := caller.Received(caller.SkipCurrent)
	return Standard("(%s): [%w]", where, who)
}
