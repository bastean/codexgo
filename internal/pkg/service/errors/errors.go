package errors

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	Bubble = errors.Bubble
	Meta   = errors.Meta
	Error  = errors.Error
)

type (
	InvalidValue = errors.InvalidValue
	AlreadyExist = errors.AlreadyExist
	NotExist     = errors.NotExist
	Failure      = errors.Failure
	Internal     = errors.Internal
)

var (
	Panic    = errors.Panic
	BubbleUp = errors.BubbleUp
)

var (
	As = errors.As
	Is = errors.Is
)

func New[Err Error](bubble *Bubble) *Err {
	return errors.New[Err](bubble)
}

func IsNot(err error, target error) bool {
	return err != nil && !Is(err, target)
}
