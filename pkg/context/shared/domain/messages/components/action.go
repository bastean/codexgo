package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	ActionMinCharactersLength = "1"
	ActionMaxCharactersLength = "20"
)

type Action struct {
	Value string `validate:"gte=1,lte=20"`
}

func NewAction(value string) (*Action, error) {
	value = strings.TrimSpace(value)

	object := &Action{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewAction",
			What:  fmt.Sprintf("Action must be between %s to %s characters", ActionMinCharactersLength, ActionMaxCharactersLength),
			Why: errors.Meta{
				"Action": value,
			},
		})
	}

	return object, nil
}
