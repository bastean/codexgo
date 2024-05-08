package valueObject

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

func NewId(id string) (model.ValueObject[string], error) {
	return sharedValueObject.NewId(id)
}
