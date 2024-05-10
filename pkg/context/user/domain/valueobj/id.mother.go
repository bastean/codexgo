package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

func RandomId() (smodel.ValueObject[string], error) {
	return NewId(sservice.Create.UUID())
}

func InvalidId() (string, error) {
	value := "x"

	_, err := NewId(value)

	return value, err
}
