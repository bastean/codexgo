package valueObjectMother

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/service/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

func RandomId() *valueObject.Id {
	return valueObject.NewId(mother.Create.UUID())
}

func InvalidId() *valueObject.Id {
	return valueObject.NewId("x")
}

func EmptyId() *valueObject.Id {
	return valueObject.NewId("")
}
