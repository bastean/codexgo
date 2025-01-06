package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

var (
	StatusOneOf = []string{"Queued", "Succeeded", "Failed", "Done"}
)

type Status struct {
	Value string `validate:"oneof=queued succeeded failed done"`
}

func NewStatus(value string) (*Status, error) {
	value = strings.TrimSpace(value)

	value = strings.ToLower(value)

	object := &Status{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewStatus",
			What:  fmt.Sprintf("Status must be only one of these values: %s", strings.Join(StatusOneOf, ", ")),
			Why: errors.Meta{
				"Status": value,
			},
		})
	}

	return object, nil
}
