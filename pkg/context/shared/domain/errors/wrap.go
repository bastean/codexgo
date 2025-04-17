package errors

import (
	"errors"
)

var (
	Standard = errors.New
	Join     = errors.Join
	As       = errors.As
	Is       = errors.Is
)
