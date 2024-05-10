package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

func RandomPassword() (smodel.ValueObject[string], error) {
	return NewPassword(sservice.Create.Regex(`[\W\w]{8,64}`))
}

func WithInvalidPasswordLength() (string, error) {
	value := "x"

	_, err := NewPassword(value)

	return value, err
}
