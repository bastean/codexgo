package user

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	UsernameMinCharactersLength = "2"
	UsernameMaxCharactersLength = "20"
)

type Username struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

func NewUsername(value string) (*Username, error) {
	value = strings.TrimSpace(value)

	object := &Username{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewUsername",
			What:  fmt.Sprintf("Username must be between %s to %s characters and be alphanumeric only", UsernameMinCharactersLength, UsernameMaxCharactersLength),
			Why: errors.Meta{
				"Username": value,
			},
		})
	}

	return object, nil
}
