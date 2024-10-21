package errs

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func Assertion(where string) error {
	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Failure in type assertion",
	})
}
