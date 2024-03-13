package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	"github.com/bastean/codexgo/test/pkg/context/shared/domain/service"
)

func NewPassword(password string) *valueObject.Password {
	return valueObject.NewPassword(password)
}

func RandomPassword() *valueObject.Password {
	return NewPassword(service.Mother.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() *valueObject.Password {
	return NewPassword("x")
}

func EmptyPassword() *valueObject.Password {
	return NewPassword("")
}
