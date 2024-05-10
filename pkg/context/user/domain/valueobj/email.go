package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/svalueobj"
)

func NewEmail(email string) (smodel.ValueObject[string], error) {
	return svalueobj.NewEmail(email)
}
