package user

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	PlainPasswordMinCharactersLength = "8"
	PlainPasswordMaxCharactersLength = "64"
)

type PlainPassword struct {
	Value string `validate:"gte=8,lte=64"`
}

type CipherPassword struct {
	Value string `validate:"required"`
}

func NewPlainPassword(value string) (*PlainPassword, error) {
	valueObj := &PlainPassword{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewPlainPassword",
			What:  fmt.Sprintf("Password must be between %s to %s characters", PlainPasswordMinCharactersLength, PlainPasswordMaxCharactersLength),
			Why: errors.Meta{
				"Password": value,
			},
		})
	}

	return valueObj, nil
}

func NewCipherPassword(value string) (*CipherPassword, error) {
	valueObj := &CipherPassword{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "NewCipherPassword",
			What:  "Cipher Password is required",
			Why: errors.Meta{
				"Password": value,
			},
		})
	}

	return valueObj, nil
}
