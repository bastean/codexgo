package errors

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
)

type (
	InvalidValue = errors.InvalidValue
	AlreadyExist = errors.AlreadyExist
	NotExist     = errors.NotExist
	Failure      = errors.Failure
	Internal     = errors.Internal
)

var (
	BubbleUp = errors.BubbleUp
	As       = errors.As
	Is       = errors.Is
)

func NewInternal(where, what string, who error) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  what,
		Who:   who,
	})
}

func NewFailure(where, what string, who error) error {
	return errors.NewFailure(&errors.Bubble{
		Where: where,
		What:  what,
		Who:   who,
	})
}
