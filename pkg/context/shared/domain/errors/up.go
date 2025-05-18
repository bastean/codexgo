package errors

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func BubbleUp(who error) error {
	_, _, where := caller.Received(caller.SkipCurrent)
	return fmt.Errorf("(%s): [%w]", where, who)
}
