package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomEmail() (model.ValueObject[string], error) {
	return valueObject.NewEmail(mother.Create.Email())
}

func InvalidEmail() (string, error) {
	value := "x"

	_, err := valueObject.NewEmail(value)

	return value, err
}
