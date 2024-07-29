package user

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type Verified struct {
	Value bool
}

func NewVerified(value bool) (*Verified, error) {
	valueObj := &Verified{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewVerified",
			What:  "Invalid verified value",
			Why: errors.Meta{
				"Verified": fmt.Sprintf("%t", value),
			},
		})
	}

	return valueObj, nil
}
