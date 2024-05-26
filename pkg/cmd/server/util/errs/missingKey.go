package errs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
)

func MissingKey(what, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "failure to get " + what + " from context",
	})
}
