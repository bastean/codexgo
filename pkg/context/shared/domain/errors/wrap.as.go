package errors

import (
	"errors"
)

func As(err error, target any) bool {
	return errors.As(err, target)
}
