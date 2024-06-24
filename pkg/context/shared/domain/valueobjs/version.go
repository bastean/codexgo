package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

type Version struct {
	Version string `validate:"number"`
}

func (value *Version) Value() string {
	return value.Version
}

func (value *Version) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewVersion(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Version{
		Version: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewVersion",
			What:  "version must be numeric only",
			Why: errors.Meta{
				"Version": value,
			},
		})
	}

	return valueObj, nil
}
