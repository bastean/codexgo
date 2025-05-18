package errs

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func MissingKey(what string) error {
	_, _, where := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  fmt.Sprintf("Failure to obtain the value of the key [%s]", what),
	})
}
