package errors

import (
	"fmt"
)

func BubbleUp(who error, where string) error {
	return fmt.Errorf("(%v): [%w]", where, who)
}
