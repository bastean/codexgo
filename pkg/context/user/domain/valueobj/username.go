package valueobj

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const UsernameMinCharactersLength = "2"
const UsernameMaxCharactersLength = "20"

type Username struct {
	Username string `validate:"gte=2,lte=20,alphanum"`
}

func (value *Username) Value() string {
	return value.Username
}

func (value *Username) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewUsername(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Username{
		Username: value,
	}

	if valueObj.IsValid() != nil {
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
