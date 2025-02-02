package errs

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func MissingKey(what, where string) error {
	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  fmt.Sprintf("Failure to obtain the value of the key [%s]", what),
	})
}
