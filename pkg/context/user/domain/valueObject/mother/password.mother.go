package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomPassword() (model.ValueObject[string], error) {
	return valueObject.NewPassword(mother.Create.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() (model.ValueObject[string], error) {
	return valueObject.NewPassword("x")
}

func EmptyPassword() (model.ValueObject[string], error) {
	return valueObject.NewPassword("")
}
