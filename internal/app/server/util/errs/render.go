package errs

import (
	"github.com/bastean/codexgo/internal/pkg/service/errors"
)

func Render(who error, where string) error {
	return errors.NewInternal(&errors.Bubble{
		Where: where,
		What:  "Cannot render a page",
		Who:   who,
	})
}
