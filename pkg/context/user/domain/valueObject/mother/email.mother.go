package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomEmail() *valueObject.Email {
	return valueObject.NewEmail(mother.Create.Email())
}

func InvalidEmail() *valueObject.Email {
	return valueObject.NewEmail("x")
}

func EmptyEmail() *valueObject.Email {
	return valueObject.NewEmail("")
}
