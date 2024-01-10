package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewEmail(email string) *valueObjects.Email {
	return valueObjects.NewEmail(email)
}

func RandomEmail() *valueObjects.Email {
	return NewEmail(mother.Creator.Email())
}

func InvalidEmail() *valueObjects.Email {
	return NewEmail("x")
}

func EmptyEmail() *valueObjects.Email {
	return NewEmail("")
}
