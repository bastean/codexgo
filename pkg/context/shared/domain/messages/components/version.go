package components

import (
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type Version struct {
	Value string `validate:"number"`
}

func NewVersion(value string) (*Version, error) {
	value = strings.TrimSpace(value)

	object := &Version{
		Value: value,
	}

	if services.IsValueObjectInvalid(object) {
		return nil, errors.New[errors.InvalidValue](&errors.Bubble{
			Where: "NewVersion",
			What:  "Version must be numeric only",
			Why: errors.Meta{
				"Version": value,
			},
		})
	}

	return object, nil
}
