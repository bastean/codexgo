package errors

import (
	"fmt"
)

func BubbleUp(who error, where string) error {
	return fmt.Errorf("(%s): [%w]", where, who)
}
