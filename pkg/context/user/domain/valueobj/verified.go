package valueobj

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
)

type Verified struct {
	Verified bool
}

func (verified *Verified) Value() bool {
	return verified.Verified
}

func (verified *Verified) IsValid() error {
	return nil
}

func NewVerified(verified bool) (models.ValueObject[bool], error) {
	verifiedVO := &Verified{
		Verified: verified,
	}

	if verifiedVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewVerified",
			What:  "invalid verified value",
			Why: errors.Meta{
				"Verified": fmt.Sprintf("%s", verified),
			},
		})
	}

	return verifiedVO, nil
}
