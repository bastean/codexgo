package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const CommandMinCharactersLength = "1"
const CommandMaxCharactersLength = "20"

type Command struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewCommand(value string) (*Command, error) {
	value = strings.TrimSpace(value)

	valueObj := &Command{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
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
