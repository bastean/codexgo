package errs

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func Missing(what string) error {
	where, _, _, _ := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Failure](&errors.Bubble{
		Where: where,
		What:  fmt.Sprintf("Missing %q", what),
	})
}

func MissingKey(what string) error {
	where, _, _, _ := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  fmt.Sprintf("Missing key %q", what),
	})
}

func MissingTokenSignature() error {
	where, _, _, _ := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Missing token signature",
	})
}
