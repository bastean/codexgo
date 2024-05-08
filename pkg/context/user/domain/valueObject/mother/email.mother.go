package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomEmail() (model.ValueObject[string], error) {
	return valueObject.NewEmail(mother.Create.Email())
}

func InvalidEmail() (model.ValueObject[string], error) {
	return valueObject.NewEmail("x")
}

func EmptyEmail() (model.ValueObject[string], error) {
	return valueObject.NewEmail("")
}
