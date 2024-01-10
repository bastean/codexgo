package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewPassword(password string) *valueObjects.Password {
	return valueObjects.NewPassword(password)
}

func RandomPassword() *valueObjects.Password {
	return NewPassword(mother.Creator.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() *valueObjects.Password {
	return NewPassword("x")
}

func EmptyPassword() *valueObjects.Password {
	return NewPassword("")
}
