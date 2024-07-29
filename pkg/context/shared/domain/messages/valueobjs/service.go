package valueobjs

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
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
			What:  fmt.Sprintf("Service must be between %s to %s characters and be alphanumeric only", ServiceMinCharactersLength, ServiceMaxCharactersLength),
			Why: errors.Meta{
				"Service": value,
			},
		})
	}

	return valueObj, nil
}
