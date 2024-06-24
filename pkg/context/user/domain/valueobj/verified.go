package valueobj

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
)

type Verified struct {
	Verified bool
}

func (value *Verified) Value() bool {
	return value.Verified
}

func (value *Verified) IsValid() error {
	return nil
}

func NewVerified(value bool) (models.ValueObject[bool], error) {
	valueObj := &Verified{
		Verified: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewVerified",
			What:  "invalid verified value",
			Why: errors.Meta{
				"Verified": fmt.Sprintf("%t", value),
			},
		})
	}

	return valueObj, nil
}
