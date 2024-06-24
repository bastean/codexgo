package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const ServiceMinCharactersLength = "1"
const ServiceMaxCharactersLength = "20"

type Service struct {
	Service string `validate:"gte=1,lte=20,alphanum"`
}

func (value *Service) Value() string {
	return value.Service
}

func (value *Service) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewService(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Service{
		Service: value,
	}

	if valueObj.IsValid() != nil {
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
