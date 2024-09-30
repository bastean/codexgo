package format

import (
	"fmt"
)

func ToString(value any) string {
	return fmt.Sprintf("%v", value)
}
