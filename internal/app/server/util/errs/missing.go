package errs

import (
	"fmt"

	"github.com/bastean/codexgo/internal/pkg/service/errors"
)

func MissingKey(what, where string) error {
	return errors.NewInternal(
		where,
		fmt.Sprintf("Failure to obtain the value of the key [%s]", what),
		nil,
	)
}
