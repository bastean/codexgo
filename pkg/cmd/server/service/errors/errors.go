package errors

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
)

type InvalidValue = errors.InvalidValue

type AlreadyExist = errors.AlreadyExist

type NotExist = errors.NotExist

type Failure = errors.Failure

type Internal = errors.Internal

var As = errors.As

var Is = errors.Is

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
