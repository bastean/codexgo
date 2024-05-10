package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/sservice"
)

func RandomUsername() (smodel.ValueObject[string], error) {
	return NewUsername(sservice.Create.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() (string, error) {
	value := "x"

	_, err := NewUsername(value)

	return value, err
}

func WithInvalidUsernameAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewUsername(value)

	return value, err
}
