package sharedValueObject

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	Value string `validate:"uuid4"`
}

var InvalidIdValue = errors.InvalidValue{Message: "Id Invalid"}

func ensureIsValidId(id *Id) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return validate.Struct(id)
}

func NewId(id string) *Id {
	id = strings.TrimSpace(id)

	idVO := &Id{id}

	err := ensureIsValidId(idVO)

	if err != nil {
		panic(InvalidIdValue)
	}

	return idVO
}
