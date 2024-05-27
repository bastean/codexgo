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

func NewInternal(where, what string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  what,
	})
}
