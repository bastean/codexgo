package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomUsername() (model.ValueObject[string], error) {
	return valueObject.NewUsername(mother.Create.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() (string, error) {
	value := "x"

	_, err := valueObject.NewUsername(value)

	return value, err
}

func WithInvalidUsernameAlphanumeric() (string, error) {
	value := "<></>"

	_, err := valueObject.NewUsername(value)

	return value, err
}
