package errors

import (
	"errors"
	"fmt"
)

var (
	Standard = fmt.Errorf
	Join     = errors.Join
	As       = errors.As
	Is       = errors.Is
)
