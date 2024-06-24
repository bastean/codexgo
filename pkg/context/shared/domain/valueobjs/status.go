package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

var StatusOneOf = []string{"queued", "succeeded", "failed", "done"}

type Status struct {
	Status string `validate:"oneof=queued succeeded failed done"`
}

func (value *Status) Value() string {
	return value.Status
}

func (value *Status) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewStatus(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Status{
		Status: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewStatus",
			What:  "status must be only one of these values: " + strings.Join(StatusOneOf, ", "),
			Why: errors.Meta{
				"Status": value,
			},
		})
	}

	return valueObj, nil
}
