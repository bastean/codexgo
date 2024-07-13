package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

var StatusOneOf = []string{"queued", "succeeded", "failed", "done"}

type Status struct {
	Value string `validate:"oneof=queued succeeded failed done"`
}

func NewStatus(value string) (*Status, error) {
	value = strings.TrimSpace(value)

	valueObj := &Status{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
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
