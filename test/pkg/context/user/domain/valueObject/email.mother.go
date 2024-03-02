package valueObject

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/test/pkg/context/shared/domain/mother"
)

func NewEmail(email string) *valueObject.Email {
	return valueObject.NewEmail(email)
}

func RandomEmail() *valueObject.Email {
	return NewEmail(mother.Creator.Email())
}

func InvalidEmail() *valueObject.Email {
	return NewEmail("x")
}

func EmptyEmail() *valueObject.Email {
	return NewEmail("")
}
