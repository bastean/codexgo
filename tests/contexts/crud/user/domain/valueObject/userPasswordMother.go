package valueObject

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewPassword(password string) *valueObject.Password {
	return valueObject.NewPassword(password)
}

func RandomPassword() *valueObject.Password {
	return NewPassword(mother.Creator.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() *valueObject.Password {
	return NewPassword("x")
}

func EmptyPassword() *valueObject.Password {
	return NewPassword("")
}
