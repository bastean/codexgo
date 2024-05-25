package errs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
)

func MissingKey(what, where string) error {
	return serror.NewInternal(&serror.Bubble{
		Where: where,
		What:  "failure to get " + what + " from context",
	})
}
