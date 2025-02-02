package errs

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func SessionSave(who error, where string) error {
	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Failure to save session",
		Who:   who,
	})
}
