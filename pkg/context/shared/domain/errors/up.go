package errors

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func BubbleUp(who error) error {
	where := "Unknown"

	_, _, name := caller.Received(caller.SkipCurrent)

	if name != "" {
		where = name
	}

	return fmt.Errorf("(%s): [%w]", where, who)
}
