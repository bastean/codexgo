package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	Bubble = errors.Bubble
	Meta   = errors.Meta
)

type (
	Default      = errors.Default
	Internal     = errors.Internal
	Failure      = errors.Failure
	InvalidValue = errors.InvalidValue
	AlreadyExist = errors.AlreadyExist
	NotExist     = errors.NotExist
)

var (
	Panic    = errors.Panic
	BubbleUp = errors.BubbleUp
)

var (
	As = errors.As
	Is = errors.Is
)

func New[Error ~struct{ *Bubble }](bubble *Bubble) *Error {
	return errors.New[Error](bubble)
}

func IsNot(err error, target error) bool {
	return err != nil && !Is(err, target)
}
