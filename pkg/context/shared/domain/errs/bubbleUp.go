package errs

import (
	"fmt"
)

func BubbleUp(where string, who error) error {
	return fmt.Errorf("(%s): [%w]", where, who)
}
