package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func RootWithValidValue() *Root {
	value, err := NewRoot()

	if err != nil {
		errors.Panic(err.Error(), "RootWithValidValue")
	}

	return value
}
