package format

import (
	"fmt"
)

func String(value any) string {
	return fmt.Sprintf("%v", value)
}
