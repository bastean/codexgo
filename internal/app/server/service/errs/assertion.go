package errs

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func Assertion() error {
	_, _, where := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Failure in type assertion",
	})
}
