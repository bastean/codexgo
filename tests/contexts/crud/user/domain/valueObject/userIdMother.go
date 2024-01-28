package valueObject

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewId(id string) *valueObject.Id {
	return valueObject.NewId(id)
}

func RandomId() *valueObject.Id {
	return NewId(mother.Creator.UUID())
}

func InvalidId() *valueObject.Id {
	return NewId("x")
}

func EmptyId() *valueObject.Id {
	return NewId("")
}
