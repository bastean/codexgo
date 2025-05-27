package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Verified struct {
	values.Bool
}

func (v *Verified) Validate() error {
	v.Valid()
	return nil
}
