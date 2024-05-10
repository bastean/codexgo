package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/svalueobj"
)

func NewId(id string) (smodel.ValueObject[string], error) {
	return svalueobj.NewId(id)
}
