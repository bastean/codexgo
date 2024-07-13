package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const ServiceMinCharactersLength = "1"
const ServiceMaxCharactersLength = "20"

type Service struct {
	Value string `validate:"gte=1,lte=20,alphanum"`
}

func NewService(value string) (*Service, error) {
	value = strings.TrimSpace(value)

	valueObj := &Service{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewService",
			What:  "service must be between " + ServiceMinCharactersLength + " to " + ServiceMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Service": value,
			},
		})
	}

	return valueObj, nil
}
