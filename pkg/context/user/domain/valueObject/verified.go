package valueObject

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
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

func NewVerified(verified bool) (model.ValueObject[bool], error) {
	verifiedVO := &Verified{
		Verified: verified,
	}

	if verifiedVO.IsValid() != nil {
		return nil, errs.NewInvalidValueError(&errs.Bubble{
			Where: "NewVerified",
			What:  "invalid",
			Why: errs.Meta{
				"Verified": fmt.Sprintf("%v", verified),
			},
		})
	}

	return verifiedVO, nil
}
