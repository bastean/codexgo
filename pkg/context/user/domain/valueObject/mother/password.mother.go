package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomPassword() *valueObject.Password {
	return valueObject.NewPassword(mother.Create.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() *valueObject.Password {
	return valueObject.NewPassword("x")
}

func EmptyPassword() *valueObject.Password {
	return valueObject.NewPassword("")
}
