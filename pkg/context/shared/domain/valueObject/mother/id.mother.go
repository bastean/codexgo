package sharedValueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

func RandomId() *sharedValueObject.Id {
	return sharedValueObject.NewId(mother.Create.UUID())
}

func InvalidId() *sharedValueObject.Id {
	return sharedValueObject.NewId("x")
}

func EmptyId() *sharedValueObject.Id {
	return sharedValueObject.NewId("")
}
