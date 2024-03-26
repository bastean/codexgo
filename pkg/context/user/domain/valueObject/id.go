package valueObject

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

type Id struct {
	Value string
}

func NewId(id string) *Id {
	return &Id{
		Value: sharedValueObject.NewId(id).Value,
	}
}
