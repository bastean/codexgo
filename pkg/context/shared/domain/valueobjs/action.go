package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const ActionMinCharactersLength = "1"
const ActionMaxCharactersLength = "20"

type Action struct {
	Action string `validate:"gte=1,lte=20"`
}

func (value *Action) Value() string {
	return value.Action
}

func (value *Action) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewAction(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Action{
		Action: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewAction",
			What:  "action must be between " + ActionMinCharactersLength + " to " + ActionMaxCharactersLength + " characters",
			Why: errors.Meta{
				"Action": value,
			},
		})
	}

	return valueObj, nil
}
