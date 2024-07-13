package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const EntityMinCharactersLength = "1"
const EntityMaxCharactersLength = "20"

type Entity struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewEntity(value string) (*Entity, error) {
	value = strings.TrimSpace(value)

	valueObj := &Entity{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEntity",
			What:  "entity must be between " + EntityMinCharactersLength + " to " + EntityMaxCharactersLength + " characters and be alpha only",
			Why: errors.Meta{
				"Entity": value,
			},
		})
	}

	return valueObj, nil
}
