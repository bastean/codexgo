package errs

import (
	"fmt"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func MissingKey(what, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  fmt.Sprintf("Failure to obtain the value of the key [%s]", what),
	})
}
