package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func TimeWithValidValue() *Time {
	value, err := values.New[*Time](services.TimeNow())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func TimeWithInvalidValue() (string, error) {
	var value string

	_, err := values.New[*Time](value)

	return value, err
}
