package user

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const UsernameMinCharactersLength = "2"
const UsernameMaxCharactersLength = "20"

type Username struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

func NewUsername(value string) (*Username, error) {
	value = strings.TrimSpace(value)

	valueObj := &Username{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewUsername",
			What:  "username must be between " + UsernameMinCharactersLength + " to " + UsernameMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Username": value,
			},
		})
	}

	return valueObj, nil
}
