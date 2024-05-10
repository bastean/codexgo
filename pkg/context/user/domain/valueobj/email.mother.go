package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

func RandomEmail() (smodel.ValueObject[string], error) {
	return NewEmail(sservice.Create.Email())
}

func InvalidEmail() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}
