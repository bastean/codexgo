package errors

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
)

type (
	ErrInvalidValue = errors.ErrInvalidValue
	ErrAlreadyExist = errors.ErrAlreadyExist
	ErrNotExist     = errors.ErrNotExist
	ErrFailure      = errors.ErrFailure
	ErrInternal     = errors.ErrInternal
)

var (
	BubbleUp = errors.BubbleUp
	As       = errors.As
	Is       = errors.Is
)

func NewFailure(where, what string, who error) error {
	return errors.NewFailure(&errors.Bubble{
		Where: where,
		What:  what,
		Who:   who,
	})
}

func NewInternal(where, what string, who error) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  what,
		Who:   who,
	})
}
