package valueobj

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
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

func NewVerified(verified bool) (smodel.ValueObject[bool], error) {
	verifiedVO := &Verified{
		Verified: verified,
	}

	if verifiedVO.IsValid() != nil {
		return nil, serror.NewInvalidValueError(&serror.Bubble{
			Where: "NewVerified",
			What:  "invalid",
			Why: serror.Meta{
				"Verified": fmt.Sprintf("%v", verified),
			},
		})
	}

	return verifiedVO, nil
}
