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

func (service *Service) Value() string {
	return service.Service
}

func (service *Service) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(service)
}

func NewService(service string) (models.ValueObject[string], error) {
	service = strings.TrimSpace(service)

	serviceVO := &Service{
		Service: service,
	}

	if serviceVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewService",
			What:  "service must be between " + ServiceMinCharactersLength + " to " + ServiceMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Service": service,
			},
		})
	}

	return serviceVO, nil
}
