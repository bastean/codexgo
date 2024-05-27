package errors

import (
	"errors"
)

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
