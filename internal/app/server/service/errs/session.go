package errs

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func SessionSave(who error, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "Failure to save session",
		Who:   who,
	})
}
