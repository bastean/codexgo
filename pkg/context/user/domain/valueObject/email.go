package valueObject

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

func NewEmail(email string) (model.ValueObject[string], error) {
	return sharedValueObject.NewEmail(email)
}
