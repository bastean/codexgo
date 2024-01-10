package valueObjects

import (
	"github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/test/contexts/crud/shared/domain/mother"
)

func NewId(id string) *valueObjects.Id {
	return valueObjects.NewId(id)
}

func RandomId() *valueObjects.Id {
	return NewId(mother.Creator.UUID())
}

func InvalidId() *valueObjects.Id {
	return NewId("x")
}

func EmptyId() *valueObjects.Id {
	return NewId("")
}
