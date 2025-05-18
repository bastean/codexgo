package errs

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

func SessionSave(who error) error {
	_, _, where := caller.Received(caller.SkipCurrent)

	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Failure to save session",
		Who:   who,
	})
}
