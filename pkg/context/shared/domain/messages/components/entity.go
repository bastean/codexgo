package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	EntityMinCharactersLength = "1"
	EntityMaxCharactersLength = "20"
)

type Entity struct {
	Value string `validate:"gte=1,lte=20,alpha"`
}

func NewEntity(value string) (*Entity, error) {
	value = strings.TrimSpace(value)

	object := &Entity{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewEntity",
			What:  fmt.Sprintf("Entity must be between %s to %s characters and be alpha only", EntityMinCharactersLength, EntityMaxCharactersLength),
			Why: errors.Meta{
				"Entity": value,
			},
		})
	}

	return object, nil
}
