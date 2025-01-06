package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	CommandMinCharactersLength = "1"
	CommandMaxCharactersLength = "20"
)

type Command struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewCommand(value string) (*Command, error) {
	value = strings.TrimSpace(value)

	object := &Command{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewCommand",
			What:  fmt.Sprintf("Command must be between %s to %s characters and be alpha only", CommandMinCharactersLength, CommandMaxCharactersLength),
			Why: errors.Meta{
				"Command": value,
			},
		})
	}

	return object, nil
}
