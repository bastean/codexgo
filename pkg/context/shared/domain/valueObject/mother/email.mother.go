package sharedValueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

func RandomEmail() *sharedValueObject.Email {
	return sharedValueObject.NewEmail(mother.Create.Email())
}

func InvalidEmail() *sharedValueObject.Email {
	return sharedValueObject.NewEmail("x")
}

func EmptyEmail() *sharedValueObject.Email {
	return sharedValueObject.NewEmail("")
}
