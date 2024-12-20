package errs

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
)

func Render(who error, where string) error {
	return errors.New[errors.Internal](&errors.Bubble{
		Where: where,
		What:  "Cannot render a page",
		Who:   who,
	})
}
