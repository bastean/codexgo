package valueObject

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

type Email struct {
	Value string
}

func NewEmail(email string) *Email {
	return &Email{
		Value: sharedValueObject.NewEmail(email).Value,
	}
}
