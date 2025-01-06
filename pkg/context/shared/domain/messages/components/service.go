package components

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	ServiceMinCharactersLength = "1"
	ServiceMaxCharactersLength = "20"
)

type Service struct {
	Value string `validate:"gte=1,lte=20,alphanum"`
}

func NewService(value string) (*Service, error) {
	value = strings.TrimSpace(value)

	object := &Service{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewService",
			What:  fmt.Sprintf("Service must be between %s to %s characters and be alphanumeric only", ServiceMinCharactersLength, ServiceMaxCharactersLength),
			Why: errors.Meta{
				"Service": value,
			},
		})
	}

	return object, nil
}
