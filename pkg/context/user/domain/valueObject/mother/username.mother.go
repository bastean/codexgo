package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomUsername() (model.ValueObject[string], error) {
	return valueObject.NewUsername(mother.Create.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() (model.ValueObject[string], error) {
	return valueObject.NewUsername("x")
}

func WithInvalidUsernameAlphanumeric() (model.ValueObject[string], error) {
	return valueObject.NewUsername("<></>")
}

func EmptyUsername() (model.ValueObject[string], error) {
	return valueObject.NewUsername("")
}
