package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const CommandMinCharactersLength = "1"
const CommandMaxCharactersLength = "20"

type Command struct {
	Command string `validate:"gte=1,lte=20,alpha"`
}

func (value *Command) Value() string {
	return value.Command
}

func (value *Command) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewCommand(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Command{
		Command: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewCommand",
			What:  "command must be between " + CommandMinCharactersLength + " to " + CommandMaxCharactersLength + " characters and be alpha only",
			Why: errors.Meta{
				"Command": value,
			},
		})
	}

	return valueObj, nil
}
